package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type StopMixStreamsServiceRequest struct {
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	MainDomainName string `position:"Query" name:"MainDomainName"`
	MixStreamName  string `position:"Query" name:"MixStreamName"`
	MixDomainName  string `position:"Query" name:"MixDomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	MainAppName    string `position:"Query" name:"MainAppName"`
	MixAppName     string `position:"Query" name:"MixAppName"`
	MainStreamName string `position:"Query" name:"MainStreamName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *StopMixStreamsServiceRequest) Invoke(client goaliyun.Client) (*StopMixStreamsServiceResponse, error) {
	resp := &StopMixStreamsServiceResponse{}
	err := client.Request("cdn", "StopMixStreamsService", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopMixStreamsServiceResponse struct {
	RequestId          goaliyun.String
	MixStreamsInfoList StopMixStreamsServiceMixStreamsInfoList
}

type StopMixStreamsServiceMixStreamsInfo struct {
	DomainName goaliyun.String
	AppName    goaliyun.String
	StreamName goaliyun.String
}

type StopMixStreamsServiceMixStreamsInfoList []StopMixStreamsServiceMixStreamsInfo

func (list *StopMixStreamsServiceMixStreamsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StopMixStreamsServiceMixStreamsInfo)
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
