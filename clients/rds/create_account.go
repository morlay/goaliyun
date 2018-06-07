package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateAccountRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AccountType          string `position:"Query" name:"AccountType"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateAccountRequest) Invoke(client goaliyun.Client) (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := client.Request("rds", "CreateAccount", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAccountResponse struct {
	RequestId goaliyun.String
}
