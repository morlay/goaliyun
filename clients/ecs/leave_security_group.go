package ecs

import (
	"github.com/morlay/goaliyun"
)

type LeaveSecurityGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *LeaveSecurityGroupRequest) Invoke(client goaliyun.Client) (*LeaveSecurityGroupResponse, error) {
	resp := &LeaveSecurityGroupResponse{}
	err := client.Request("ecs", "LeaveSecurityGroup", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type LeaveSecurityGroupResponse struct {
	RequestId goaliyun.String
}
