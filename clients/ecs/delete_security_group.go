package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteSecurityGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteSecurityGroupRequest) Invoke(client goaliyun.Client) (*DeleteSecurityGroupResponse, error) {
	resp := &DeleteSecurityGroupResponse{}
	err := client.Request("ecs", "DeleteSecurityGroup", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSecurityGroupResponse struct {
	RequestId goaliyun.String
}
