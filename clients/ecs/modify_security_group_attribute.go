package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifySecurityGroupAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityGroupName    string `position:"Query" name:"SecurityGroupName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityGroupAttributeRequest) Invoke(client goaliyun.Client) (*ModifySecurityGroupAttributeResponse, error) {
	resp := &ModifySecurityGroupAttributeResponse{}
	err := client.Request("ecs", "ModifySecurityGroupAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityGroupAttributeResponse struct {
	RequestId goaliyun.String
}
