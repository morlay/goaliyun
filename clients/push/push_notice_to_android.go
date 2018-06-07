package push

import (
	"github.com/morlay/goaliyun"
)

type PushNoticeToAndroidRequest struct {
	ExtParameters string `position:"Query" name:"ExtParameters"`
	AppKey        int64  `position:"Query" name:"AppKey"`
	TargetValue   string `position:"Query" name:"TargetValue"`
	Title         string `position:"Query" name:"Title"`
	Body          string `position:"Query" name:"Body"`
	JobKey        string `position:"Query" name:"JobKey"`
	Target        string `position:"Query" name:"Target"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *PushNoticeToAndroidRequest) Invoke(client goaliyun.Client) (*PushNoticeToAndroidResponse, error) {
	resp := &PushNoticeToAndroidResponse{}
	err := client.Request("push", "PushNoticeToAndroid", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushNoticeToAndroidResponse struct {
	RequestId goaliyun.String
	MessageId goaliyun.String
}
