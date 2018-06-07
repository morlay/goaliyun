package polardb

import (
	"github.com/morlay/goaliyun"
)

type ResetAccountPasswordRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ResetAccountPasswordRequest) Invoke(client goaliyun.Client) (*ResetAccountPasswordResponse, error) {
	resp := &ResetAccountPasswordResponse{}
	err := client.Request("polardb", "ResetAccountPassword", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetAccountPasswordResponse struct {
	RequestId goaliyun.String
}
