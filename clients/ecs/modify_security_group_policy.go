package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifySecurityGroupPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InnerAccessPolicy    string `position:"Query" name:"InnerAccessPolicy"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityGroupPolicyRequest) Invoke(client goaliyun.Client) (*ModifySecurityGroupPolicyResponse, error) {
	resp := &ModifySecurityGroupPolicyResponse{}
	err := client.Request("ecs", "ModifySecurityGroupPolicy", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityGroupPolicyResponse struct {
	RequestId goaliyun.String
}
