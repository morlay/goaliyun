package rds

import (
	"github.com/morlay/goaliyun"
)

type ResetAccountRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ResetAccountRequest) Invoke(client goaliyun.Client) (*ResetAccountResponse, error) {
	resp := &ResetAccountResponse{}
	err := client.Request("rds", "ResetAccount", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetAccountResponse struct {
	RequestId goaliyun.String
}
