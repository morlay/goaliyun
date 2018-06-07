package push

import (
	"github.com/morlay/goaliyun"
)

type PushNoticeToiOSRequest struct {
	ExtParameters string `position:"Query" name:"ExtParameters"`
	ApnsEnv       string `position:"Query" name:"ApnsEnv"`
	AppKey        int64  `position:"Query" name:"AppKey"`
	TargetValue   string `position:"Query" name:"TargetValue"`
	Title         string `position:"Query" name:"Title"`
	Body          string `position:"Query" name:"Body"`
	JobKey        string `position:"Query" name:"JobKey"`
	Target        string `position:"Query" name:"Target"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *PushNoticeToiOSRequest) Invoke(client goaliyun.Client) (*PushNoticeToiOSResponse, error) {
	resp := &PushNoticeToiOSResponse{}
	err := client.Request("push", "PushNoticeToiOS", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushNoticeToiOSResponse struct {
	RequestId goaliyun.String
	MessageId goaliyun.String
}
