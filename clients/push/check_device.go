package push

import (
	"github.com/morlay/goaliyun"
)

type CheckDeviceRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CheckDeviceRequest) Invoke(client goaliyun.Client) (*CheckDeviceResponse, error) {
	resp := &CheckDeviceResponse{}
	err := client.Request("push", "CheckDevice", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckDeviceResponse struct {
	RequestId goaliyun.String
	Available bool
}
