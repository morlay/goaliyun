package live

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
	err := client.Request("live", "ForbidLiveStream", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ForbidLiveStreamResponse struct {
	RequestId goaliyun.String
}
