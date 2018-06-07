package rds

import (
	"github.com/morlay/goaliyun"
)

type CheckAccountNameAvailableRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CheckAccountNameAvailableRequest) Invoke(client goaliyun.Client) (*CheckAccountNameAvailableResponse, error) {
	resp := &CheckAccountNameAvailableResponse{}
	err := client.Request("rds", "CheckAccountNameAvailable", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckAccountNameAvailableResponse struct {
	RequestId goaliyun.String
}
