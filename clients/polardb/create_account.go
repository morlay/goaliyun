package polardb

import (
	"github.com/morlay/goaliyun"
)

type CreateAccountRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DatabaseName         string `position:"Query" name:"DatabaseName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateAccountRequest) Invoke(client goaliyun.Client) (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := client.Request("polardb", "CreateAccount", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAccountResponse struct {
	RequestId goaliyun.String
}
