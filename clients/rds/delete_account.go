package rds

import (
	"github.com/morlay/goaliyun"
)

type DeleteAccountRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteAccountRequest) Invoke(client goaliyun.Client) (*DeleteAccountResponse, error) {
	resp := &DeleteAccountResponse{}
	err := client.Request("rds", "DeleteAccount", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAccountResponse struct {
	RequestId goaliyun.String
}
