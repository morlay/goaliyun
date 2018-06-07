package push

import (
	"github.com/morlay/goaliyun"
)

type UnbindPhoneRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UnbindPhoneRequest) Invoke(client goaliyun.Client) (*UnbindPhoneResponse, error) {
	resp := &UnbindPhoneResponse{}
	err := client.Request("push", "UnbindPhone", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnbindPhoneResponse struct {
	RequestId goaliyun.String
}
