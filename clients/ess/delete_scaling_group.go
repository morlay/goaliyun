package ess

import (
	"github.com/morlay/goaliyun"
)

type DeleteScalingGroupRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	ForceDelete          string `position:"Query" name:"ForceDelete"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteScalingGroupRequest) Invoke(client goaliyun.Client) (*DeleteScalingGroupResponse, error) {
	resp := &DeleteScalingGroupResponse{}
	err := client.Request("ess", "DeleteScalingGroup", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteScalingGroupResponse struct {
	RequestId goaliyun.String
}
