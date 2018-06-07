package iot

import (
	"github.com/morlay/goaliyun"
)

type EnableThingRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *EnableThingRequest) Invoke(client goaliyun.Client) (*EnableThingResponse, error) {
	resp := &EnableThingResponse{}
	err := client.Request("iot", "EnableThing", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnableThingResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
}
