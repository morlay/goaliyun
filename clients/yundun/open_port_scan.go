package yundun

import (
	"github.com/morlay/goaliyun"
)

type OpenPortScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *OpenPortScanRequest) Invoke(client goaliyun.Client) (*OpenPortScanResponse, error) {
	resp := &OpenPortScanResponse{}
	err := client.Request("yundun", "OpenPortScan", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OpenPortScanResponse struct {
	RequestId goaliyun.String
}
