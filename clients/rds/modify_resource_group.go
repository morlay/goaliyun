package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyResourceGroupRequest struct {
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyResourceGroupRequest) Invoke(client goaliyun.Client) (*ModifyResourceGroupResponse, error) {
	resp := &ModifyResourceGroupResponse{}
	err := client.Request("rds", "ModifyResourceGroup", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyResourceGroupResponse struct {
	RequestId goaliyun.String
}
