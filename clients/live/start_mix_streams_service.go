package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type StartMixStreamsServiceRequest struct {
	MixType        string `position:"Query" name:"MixType"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	MainDomainName string `position:"Query" name:"MainDomainName"`
	MixStreamName  string `position:"Query" name:"MixStreamName"`
	MixTemplate    string `position:"Query" name:"MixTemplate"`
	MixDomainName  string `position:"Query" name:"MixDomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	MainAppName    string `position:"Query" name:"MainAppName"`
	MixAppName     string `position:"Query" name:"MixAppName"`
	MainStreamName string `position:"Query" name:"MainStreamName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *StartMixStreamsServiceRequest) Invoke(client goaliyun.Client) (*StartMixStreamsServiceResponse, error) {
	resp := &StartMixStreamsServiceResponse{}
	err := client.Request("live", "StartMixStreamsService", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartMixStreamsServiceResponse struct {
	RequestId          goaliyun.String
	MixStreamsInfoList StartMixStreamsServiceMixStreamsInfoList
}

type StartMixStreamsServiceMixStreamsInfo struct {
	DomainName goaliyun.String
	AppName    goaliyun.String
	StreamName goaliyun.String
}

type StartMixStreamsServiceMixStreamsInfoList []StartMixStreamsServiceMixStreamsInfo

func (list *StartMixStreamsServiceMixStreamsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartMixStreamsServiceMixStreamsInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
