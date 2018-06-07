package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetHttpHeaderConfigRequest struct {
	HeaderValue   string `position:"Query" name:"HeaderValue"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      int64  `position:"Query" name:"ConfigId"`
	DomainName    string `position:"Query" name:"DomainName"`
	HeaderKey     string `position:"Query" name:"HeaderKey"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetHttpHeaderConfigRequest) Invoke(client goaliyun.Client) (*SetHttpHeaderConfigResponse, error) {
	resp := &SetHttpHeaderConfigResponse{}
	err := client.Request("cdn", "SetHttpHeaderConfig", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetHttpHeaderConfigResponse struct {
	RequestId goaliyun.String
}
