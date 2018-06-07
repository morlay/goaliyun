package yundun

import (
	"github.com/morlay/goaliyun"
)

type OpenVulScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *OpenVulScanRequest) Invoke(client goaliyun.Client) (*OpenVulScanResponse, error) {
	resp := &OpenVulScanResponse{}
	err := client.Request("yundun", "OpenVulScan", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OpenVulScanResponse struct {
	RequestId goaliyun.String
}
