package iot

import (
	"github.com/morlay/goaliyun"
)

type GetDeviceStatusRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetDeviceStatusRequest) Invoke(client goaliyun.Client) (*GetDeviceStatusResponse, error) {
	resp := &GetDeviceStatusResponse{}
	err := client.Request("iot", "GetDeviceStatus", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetDeviceStatusResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         GetDeviceStatusData
}

type GetDeviceStatusData struct {
	Status goaliyun.String
}
