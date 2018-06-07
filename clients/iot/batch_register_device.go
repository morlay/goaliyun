package iot

import (
	"github.com/morlay/goaliyun"
)

type BatchRegisterDeviceRequest struct {
	Count      int64  `position:"Query" name:"Count"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *BatchRegisterDeviceRequest) Invoke(client goaliyun.Client) (*BatchRegisterDeviceResponse, error) {
	resp := &BatchRegisterDeviceResponse{}
	err := client.Request("iot", "BatchRegisterDevice", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchRegisterDeviceResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         BatchRegisterDeviceData
}

type BatchRegisterDeviceData struct {
	ApplyId goaliyun.Integer
}
