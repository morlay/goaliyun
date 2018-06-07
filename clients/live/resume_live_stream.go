package live

import (
	"github.com/morlay/goaliyun"
)

type ResumeLiveStreamRequest struct {
	AppName        string `position:"Query" name:"AppName"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	LiveStreamType string `position:"Query" name:"LiveStreamType"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	StreamName     string `position:"Query" name:"StreamName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ResumeLiveStreamRequest) Invoke(client goaliyun.Client) (*ResumeLiveStreamResponse, error) {
	resp := &ResumeLiveStreamResponse{}
	err := client.Request("live", "ResumeLiveStream", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResumeLiveStreamResponse struct {
	RequestId goaliyun.String
}
