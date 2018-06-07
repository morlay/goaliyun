package iot

import (
	"github.com/morlay/goaliyun"
)

type DisableThingRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DisableThingRequest) Invoke(client goaliyun.Client) (*DisableThingResponse, error) {
	resp := &DisableThingResponse{}
	err := client.Request("iot", "DisableThing", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DisableThingResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
}
