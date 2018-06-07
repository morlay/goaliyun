package push

import (
	"github.com/morlay/goaliyun"
)

type PushMessageToiOSRequest struct {
	AppKey      int64  `position:"Query" name:"AppKey"`
	TargetValue string `position:"Query" name:"TargetValue"`
	Title       string `position:"Query" name:"Title"`
	Body        string `position:"Query" name:"Body"`
	JobKey      string `position:"Query" name:"JobKey"`
	Target      string `position:"Query" name:"Target"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *PushMessageToiOSRequest) Invoke(client goaliyun.Client) (*PushMessageToiOSResponse, error) {
	resp := &PushMessageToiOSResponse{}
	err := client.Request("push", "PushMessageToiOS", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PushMessageToiOSResponse struct {
	RequestId goaliyun.String
	MessageId goaliyun.String
}
