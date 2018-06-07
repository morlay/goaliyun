package drds

import (
	"github.com/morlay/goaliyun"
)

type ModifyFullTableScanRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	TableNames     string `position:"Query" name:"TableNames"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	FullTableScan  string `position:"Query" name:"FullTableScan"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ModifyFullTableScanRequest) Invoke(client goaliyun.Client) (*ModifyFullTableScanResponse, error) {
	resp := &ModifyFullTableScanResponse{}
	err := client.Request("drds", "ModifyFullTableScan", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFullTableScanResponse struct {
	RequestId goaliyun.String
	Success   bool
}
