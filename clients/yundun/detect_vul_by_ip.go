package yundun

import (
	"github.com/morlay/goaliyun"
)

type DetectVulByIpRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	VulIp      string `position:"Query" name:"VulIp"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DetectVulByIpRequest) Invoke(client goaliyun.Client) (*DetectVulByIpResponse, error) {
	resp := &DetectVulByIpResponse{}
	err := client.Request("yundun", "DetectVulByIp", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetectVulByIpResponse struct {
	RequestId goaliyun.String
}
