package iot

import (
	"github.com/morlay/goaliyun"
)

type RemoveThingTopoRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RemoveThingTopoRequest) Invoke(client goaliyun.Client) (*RemoveThingTopoResponse, error) {
	resp := &RemoveThingTopoResponse{}
	err := client.Request("iot", "RemoveThingTopo", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveThingTopoResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         bool
}
