package goaliyun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/morlay/goaliyun/signature"
	"github.com/morlay/goaliyun/transform"
)

type Client interface {
	Request(service string, action string, version string, regionId string, method string, urlPattern string) RequestSender
}

type RequestSender interface {
	Do(req interface{}, resp interface{}) error
}

type MaybeErrorResponse interface {
	GetError() error
}

func NewClientWithCredential(accessKey, secretKey string, opts ...*ClientOption) *ClientWithCredential {
	return &ClientWithCredential{
		ClientOption: ComposeClientOptions(opts...),
		Credential: signature.Credential{
			AccessKey: accessKey,
			SecretKey: secretKey,
		},
	}
}

func ComposeClientOptions(opts ...*ClientOption) *ClientOption {
	if len(opts) == 0 {
		return nil
	}

	opt := &ClientOption{}

	for i := range opts {
		o := opts[i]
		if o.Timeout != 0 {
			opt.Timeout = o.Timeout
		}
		if o.Transports != nil {
			opt.Transports = o.Transports
		}
	}

	return opt
}

func ClientOptionWithTimeout(timeout time.Duration) *ClientOption {
	return &ClientOption{
		Timeout: timeout,
	}
}

func ClientOptionWithTransports(transports ...transform.Transport) *ClientOption {
	return &ClientOption{
		Transports: transports,
	}
}

type ClientOption struct {
	Timeout    time.Duration
	Transports []transform.Transport
}

type ClientWithCredential struct {
	*ClientOption
	signature.Credential
}

type DescribeEndpointsRequest struct {
	RegionId    string `name:"Id" position:"Query"`
	ServiceCode string `name:"ServiceCode" position:"Query"`
	Type        string `name:"Type" position:"Query"`
}

type Endpoint struct {
	Type        string
	Namespace   string
	RegionId    string `json:"Id"`
	ServiceCode string `json:"SerivceCode"`
	Endpoint    string `json:"Endpoint"`
	Protocols struct {
		Protocols []string
	}
}

func (e *Endpoint) String() string {
	return e.Protocols.Protocols[0] + "://" + e.Endpoint
}

var Endpoints = &EndpointCollection{}

type EndpointCollection struct {
	m sync.Map
}

func (ec *EndpointCollection) Register(regionId string, serviceCode string, e *Endpoint) {
	ec.m.Store(regionId+serviceCode, e)
}

func (ec *EndpointCollection) Get(regionId string, serviceCode string) *Endpoint {
	if v, ok := ec.m.Load(regionId + serviceCode); ok {
		return v.(*Endpoint)
	}
	return nil
}

func (c *ClientWithCredential) resolveEndpoint(regionId string, serviceCode string) *Endpoint {
	e := Endpoints.Get(regionId, serviceCode)
	if e != nil {
		return e
	}

	var v struct {
		Endpoints struct {
			Endpoints []Endpoint `json:"Endpoint"`
		}
	}

	err := c.Request("location", "DescribeEndpoints", "2015-06-12", "", "", "").
		Do(&DescribeEndpointsRequest{
		RegionId:    regionId,
		ServiceCode: serviceCode,
		Type:        "openAPI",
	}, &v)

	if err != nil {
		return nil
	}

	for i := range v.Endpoints.Endpoints {
		e := &v.Endpoints.Endpoints[i]
		Endpoints.Register(regionId, serviceCode, e)
		return e
	}

	return nil
}

func (c *ClientWithCredential) Request(serviceCode string, action string, version string, regionId string, method string, urlPattern string) RequestSender {
	opt := c.ClientOption
	if opt == nil {
		opt = &ClientOption{}
	}

	opt.Transports = append(opt.Transports, signature.NewAutoSignTransport(&c.Credential))

	origin := fmt.Sprintf("https://%s.aliyuncs.com", serviceCode)

	if serviceCode != "location" {
		endpoint := c.resolveEndpoint(regionId, serviceCode)
		if endpoint != nil {
			origin = endpoint.String()
		}
	}

	return &httpRequestSender{
		serviceCode:  serviceCode,
		action:       action,
		version:      version,
		method:       method,
		origin:       origin,
		urlPattern:   urlPattern,
		ClientOption: *opt,
	}
}

type httpRequestSender struct {
	serviceCode string
	version     string
	action      string
	method      string
	origin      string
	urlPattern  string
	ClientOption
}

func (r *httpRequestSender) Do(req interface{}, resp interface{}) error {
	request, errForTransform := r.transformRequest(r.method, r.origin, r.urlPattern, req)
	if errForTransform != nil {
		return NewAliyunError(
			"SDK.TransformRequestFailed",
			errForTransform.Error(),
		)
	}

	if r.Timeout == 0 {
		r.Timeout = 60 * time.Second
	}

	client := &http.Client{}
	client.Timeout = r.Timeout
	client.Transport = transform.ComposeTransports(r.Transports...)(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   client.Timeout,
			KeepAlive: 0,
		}).DialContext,
		DisableKeepAlives: true,
	})

	response, errForDo := client.Do(request)
	if errForDo != nil {
		return NewAliyunError(
			"SDK.DoRequestFailed",
			errForDo.Error(),
		)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return NewAliyunError(
			"SDK.ResponseBodyReadFailed",
			err.Error(),
		)
	}

	if response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusNetworkAuthenticationRequired {
		aliyunErr := &AliyunError{}
		if err = json.Unmarshal(data, aliyunErr); err != nil {
			return NewAliyunError(
				"SDK.JSONUnmarshalFailed",
				err.Error(),
			)
		}
		return aliyunErr
	}

	if err := json.Unmarshal(data, resp); err != nil {
		return NewAliyunError(
			"SDK.JSONUnmarshalFailed",
			err.Error(),
		)
	}

	return nil
}

func (r *httpRequestSender) transformRequest(method, origin, urlPattern string, req interface{}) (*http.Request, error) {
	if method == "" {
		method = "GET"
	}

	u, _ := url.Parse(origin)

	s := transform.NewParameterScanner()
	if err := s.Scan(reflect.ValueOf(req), ""); err != nil {
		return nil, err
	}

	params := s.Params()

	params.Set("Format", "JSON")
	params.Set("Action", r.action)
	params.Set("Version", r.version)

	if urlPattern != "" {
		path := urlPattern
		pathParams := s.PathParams()
		for key := range pathParams {
			path = strings.Replace(path, "["+key+"]", pathParams.Get(key), 1)
		}
		u.Path = path
	}

	u.RawQuery = params.Encode()

	request, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if urlPattern != "" {
		// roa way
		request.Header.Set("x-acs-version", r.version)
	}

	return request, nil
}

type TencentCloudResponse struct {
	Response interface{} `json:"Response"`
}

type TencentCloudBaseResponse struct {
	RequestId string       `json:"RequestId"`
	Error     *AliyunError `json:"Error,omitempty"`
}

func (resp *TencentCloudBaseResponse) GetError() error {
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func NewAliyunError(code string, msg string) *AliyunError {
	return &AliyunError{
		Code:    code,
		Message: msg,
	}
}

type AliyunError struct {
	RequestId string
	Message   string
	Recommend string
	HostId    string
	Code      string
}

func (e *AliyunError) Error() string {
	return fmt.Sprintf("[%s][%s] %s", e.Code, e.Message, e.HostId)
}

func NewLogTransport() transform.Transport {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &LogTransport{
			NextRoundTripper: rt,
		}
	}
}

func NewLogTransportWithBody() transform.Transport {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &LogTransport{
			WithBody:         true,
			NextRoundTripper: rt,
		}
	}
}

type LogTransport struct {
	WithBody         bool
	NextRoundTripper http.RoundTripper
}

func (t *LogTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	buf := &bytes.Buffer{}

	for k, vs := range request.Header {
		for _, v := range vs {
			if k == "" {
				continue
			}
			buf.WriteString(k)
			buf.WriteString(": ")
			buf.WriteString(v)
			buf.WriteRune('\n')
		}
	}

	resp, err := t.NextRoundTripper.RoundTrip(request)

	if t.WithBody {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if err := resp.Body.Close(); err != nil {
			return nil, err
		}
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		fmt.Printf(`%s %s
%s

========================================================

%s
`,
			request.Method,
			request.URL.String(),
			buf.String(),
			string(bodyBytes),
		)
	} else {
		fmt.Printf(`%s %s
	%s
	`,
			request.Method,
			request.URL.String(),
			buf.String(),
		)
	}

	return resp, err
}
