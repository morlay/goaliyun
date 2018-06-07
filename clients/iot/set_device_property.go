package iot

import (
	"github.com/morlay/goaliyun"
)

type SetDevicePropertyRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	Items      string `position:"Query" name:"Items"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *SetDevicePropertyRequest) Invoke(client goaliyun.Client) (*SetDevicePropertyResponse, error) {
	resp := &SetDevicePropertyResponse{}
	err := client.Request("iot", "SetDeviceProperty", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDevicePropertyResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
}
