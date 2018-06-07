package iot

import (
	"github.com/morlay/goaliyun"
)

type DeleteDeviceRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteDeviceRequest) Invoke(client goaliyun.Client) (*DeleteDeviceResponse, error) {
	resp := &DeleteDeviceResponse{}
	err := client.Request("iot", "DeleteDevice", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDeviceResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
}
