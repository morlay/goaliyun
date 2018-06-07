package iot

import (
	"github.com/morlay/goaliyun"
)

type InvokeThingServiceRequest struct {
	Args       string `position:"Query" name:"Args"`
	Identifier string `position:"Query" name:"Identifier"`
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *InvokeThingServiceRequest) Invoke(client goaliyun.Client) (*InvokeThingServiceResponse, error) {
	resp := &InvokeThingServiceResponse{}
	err := client.Request("iot", "InvokeThingService", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InvokeThingServiceResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         InvokeThingServiceData
}

type InvokeThingServiceData struct {
	Result goaliyun.String
}
