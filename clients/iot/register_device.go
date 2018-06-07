package iot

import (
	"github.com/morlay/goaliyun"
)

type RegisterDeviceRequest struct {
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RegisterDeviceRequest) Invoke(client goaliyun.Client) (*RegisterDeviceResponse, error) {
	resp := &RegisterDeviceResponse{}
	err := client.Request("iot", "RegisterDevice", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RegisterDeviceResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         RegisterDeviceData
}

type RegisterDeviceData struct {
	IotId        goaliyun.String
	ProductKey   goaliyun.String
	DeviceName   goaliyun.String
	DeviceSecret goaliyun.String
}
