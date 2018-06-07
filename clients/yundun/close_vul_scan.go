package yundun

import (
	"github.com/morlay/goaliyun"
)

type CloseVulScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *CloseVulScanRequest) Invoke(client goaliyun.Client) (*CloseVulScanResponse, error) {
	resp := &CloseVulScanResponse{}
	err := client.Request("yundun", "CloseVulScan", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CloseVulScanResponse struct {
	RequestId goaliyun.String
}
