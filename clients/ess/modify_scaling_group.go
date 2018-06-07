package ess

import (
	"github.com/morlay/goaliyun"
)

type ModifyScalingGroupRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName             string `position:"Query" name:"ScalingGroupName"`
	ScalingGroupId               string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	ActiveScalingConfigurationId string `position:"Query" name:"ActiveScalingConfigurationId"`
	MinSize                      int64  `position:"Query" name:"MinSize"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	MaxSize                      int64  `position:"Query" name:"MaxSize"`
	DefaultCooldown              int64  `position:"Query" name:"DefaultCooldown"`
	RemovalPolicy1               string `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2               string `position:"Query" name:"RemovalPolicy.2"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *ModifyScalingGroupRequest) Invoke(client goaliyun.Client) (*ModifyScalingGroupResponse, error) {
	resp := &ModifyScalingGroupResponse{}
	err := client.Request("ess", "ModifyScalingGroup", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyScalingGroupResponse struct {
	RequestId goaliyun.String
}
