package yundun

import (
	"github.com/morlay/goaliyun"
)

type DetectVulByIdRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	VulId      int64  `position:"Query" name:"VulId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DetectVulByIdRequest) Invoke(client goaliyun.Client) (*DetectVulByIdResponse, error) {
	resp := &DetectVulByIdResponse{}
	err := client.Request("yundun", "DetectVulById", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetectVulByIdResponse struct {
	RequestId goaliyun.String
}
