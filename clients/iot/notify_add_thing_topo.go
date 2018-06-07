package iot

import (
	"github.com/morlay/goaliyun"
)

type NotifyAddThingTopoRequest struct {
	GwProductKey  string `position:"Query" name:"GwProductKey"`
	GwDeviceName  string `position:"Query" name:"GwDeviceName"`
	GwIotId       string `position:"Query" name:"GwIotId"`
	DeviceListStr string `position:"Query" name:"DeviceListStr"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *NotifyAddThingTopoRequest) Invoke(client goaliyun.Client) (*NotifyAddThingTopoResponse, error) {
	resp := &NotifyAddThingTopoResponse{}
	err := client.Request("iot", "NotifyAddThingTopo", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type NotifyAddThingTopoResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
}
