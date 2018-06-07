package push

import (
	"github.com/morlay/goaliyun"
)

type PushMessageToAndroidRequest struct {
	AppKey      int64  `position:"Query" name:"AppKey"`
	TargetValue string `position:"Query" name:"TargetValue"`
	Title       string `position:"Query" name:"Title"`
	Body        string `position:"Query" name:"Body"`
	JobKey      string `position:"Query" name:"JobKey"`
	Target      string `position:"Query" name:"Target"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *PushMessageToAndroidRequest) Invoke(client goaliyun.Client) (*PushMessageToAndroidResponse, error) {
	resp := &PushMessageToAndroidResponse{}
	err := client.Request("push", "PushMessageToAndroid", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushMessageToAndroidResponse struct {
	RequestId goaliyun.String
	MessageId goaliyun.String
}
