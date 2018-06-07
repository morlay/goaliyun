package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyAccountDescriptionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyAccountDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyAccountDescriptionResponse, error) {
	resp := &ModifyAccountDescriptionResponse{}
	err := client.Request("rds", "ModifyAccountDescription", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAccountDescriptionResponse struct {
	RequestId goaliyun.String
}
