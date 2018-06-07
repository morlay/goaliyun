package ess

import (
	"github.com/morlay/goaliyun"
)

type DescribeLimitationRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeLimitationRequest) Invoke(client goaliyun.Client) (*DescribeLimitationResponse, error) {
	resp := &DescribeLimitationResponse{}
	err := client.Request("ess", "DescribeLimitation", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLimitationResponse struct {
	MaxNumberOfScalingGroups         goaliyun.Integer
	MaxNumberOfScalingConfigurations goaliyun.Integer
	MaxNumberOfScalingRules          goaliyun.Integer
	MaxNumberOfScheduledTasks        goaliyun.Integer
	MaxNumberOfScalingInstances      goaliyun.Integer
	MaxNumberOfDBInstances           goaliyun.Integer
	MaxNumberOfLoadBalancers         goaliyun.Integer
	MaxNumberOfMinSize               goaliyun.Integer
	MaxNumberOfMaxSize               goaliyun.Integer
}
