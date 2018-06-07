package polardb

import (
	"github.com/morlay/goaliyun"
)

type ModifyAccountDescriptionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyAccountDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyAccountDescriptionResponse, error) {
	resp := &ModifyAccountDescriptionResponse{}
	err := client.Request("polardb", "ModifyAccountDescription", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAccountDescriptionResponse struct {
	RequestId goaliyun.String
}
