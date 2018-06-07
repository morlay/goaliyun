package ess

import (
	"github.com/morlay/goaliyun"
)

type DisableScalingGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DisableScalingGroupRequest) Invoke(client goaliyun.Client) (*DisableScalingGroupResponse, error) {
	resp := &DisableScalingGroupResponse{}
	err := client.Request("ess", "DisableScalingGroup", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DisableScalingGroupResponse struct {
	RequestId goaliyun.String
}
