package cdn

import (
	"github.com/morlay/goaliyun"
)

type ForbidLiveStreamRequest struct {
	ResumeTime     string `position:"Query" name:"ResumeTime"`
	AppName        string `position:"Query" name:"AppName"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	LiveStreamType string `position:"Query" name:"LiveStreamType"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	StreamName     string `position:"Query" name:"StreamName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ForbidLiveStreamRequest) Invoke(client goaliyun.Client) (*ForbidLiveStreamResponse, error) {
	resp := &ForbidLiveStreamResponse{}
	err := client.Request("cdn", "ForbidLiveStream", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ForbidLiveStreamResponse struct {
	RequestId goaliyun.String
}
