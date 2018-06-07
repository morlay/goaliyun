package rds

import (
	"github.com/morlay/goaliyun"
)

type ResetAccountForPGRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ResetAccountForPGRequest) Invoke(client goaliyun.Client) (*ResetAccountForPGResponse, error) {
	resp := &ResetAccountForPGResponse{}
	err := client.Request("rds", "ResetAccountForPG", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetAccountForPGResponse struct {
	RequestId goaliyun.String
}
