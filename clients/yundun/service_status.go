package yundun

import (
	"github.com/morlay/goaliyun"
)

type ServiceStatusRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ServiceStatusRequest) Invoke(client goaliyun.Client) (*ServiceStatusResponse, error) {
	resp := &ServiceStatusResponse{}
	err := client.Request("yundun", "ServiceStatus", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ServiceStatusResponse struct {
	RequestId goaliyun.String
	PortScan  bool
	VulScan   bool
}
