package ram

import (
	"github.com/morlay/goaliyun"
)

type UnbindMFADeviceRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UnbindMFADeviceRequest) Invoke(client goaliyun.Client) (*UnbindMFADeviceResponse, error) {
	resp := &UnbindMFADeviceResponse{}
	err := client.Request("ram", "UnbindMFADevice", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnbindMFADeviceResponse struct {
	RequestId goaliyun.String
	MFADevice UnbindMFADeviceMFADevice
}

type UnbindMFADeviceMFADevice struct {
	SerialNumber goaliyun.String
}
