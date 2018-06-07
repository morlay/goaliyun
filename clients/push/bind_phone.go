package push

import (
	"github.com/morlay/goaliyun"
)

type BindPhoneRequest struct {
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	AppKey      int64  `position:"Query" name:"AppKey"`
	DeviceId    string `position:"Query" name:"DeviceId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *BindPhoneRequest) Invoke(client goaliyun.Client) (*BindPhoneResponse, error) {
	resp := &BindPhoneResponse{}
	err := client.Request("push", "BindPhone", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BindPhoneResponse struct {
	RequestId goaliyun.String
}
