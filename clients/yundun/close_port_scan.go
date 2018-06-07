package yundun

import (
	"github.com/morlay/goaliyun"
)

type ClosePortScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ClosePortScanRequest) Invoke(client goaliyun.Client) (*ClosePortScanResponse, error) {
	resp := &ClosePortScanResponse{}
	err := client.Request("yundun", "ClosePortScan", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ClosePortScanResponse struct {
	RequestId goaliyun.String
}
