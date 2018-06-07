package push

import (
	"github.com/morlay/goaliyun"
)

type CancelPushRequest struct {
	MessageId int64  `position:"Query" name:"MessageId"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *CancelPushRequest) Invoke(client goaliyun.Client) (*CancelPushResponse, error) {
	resp := &CancelPushResponse{}
	err := client.Request("push", "CancelPush", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelPushResponse struct {
	RequestId goaliyun.String
}
