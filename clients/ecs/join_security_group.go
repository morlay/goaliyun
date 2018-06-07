package ecs

import (
	"github.com/morlay/goaliyun"
)

type JoinSecurityGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *JoinSecurityGroupRequest) Invoke(client goaliyun.Client) (*JoinSecurityGroupResponse, error) {
	resp := &JoinSecurityGroupResponse{}
	err := client.Request("ecs", "JoinSecurityGroup", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type JoinSecurityGroupResponse struct {
	RequestId goaliyun.String
}
