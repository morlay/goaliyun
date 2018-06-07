package drds

import (
	"github.com/morlay/goaliyun"
)

type DescribeReadOnlyAccountRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeReadOnlyAccountRequest) Invoke(client goaliyun.Client) (*DescribeReadOnlyAccountResponse, error) {
	resp := &DescribeReadOnlyAccountResponse{}
	err := client.Request("drds", "DescribeReadOnlyAccount", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeReadOnlyAccountResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeReadOnlyAccountData
}

type DescribeReadOnlyAccountData struct {
	DbName         goaliyun.String
	DrdsInstanceId goaliyun.String
	AccountName    goaliyun.String
}
