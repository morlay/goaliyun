package ess

import (
	"github.com/morlay/goaliyun"
)

type ExecuteScalingRuleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingRuleAri       string `position:"Query" name:"ScalingRuleAri"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ExecuteScalingRuleRequest) Invoke(client goaliyun.Client) (*ExecuteScalingRuleResponse, error) {
	resp := &ExecuteScalingRuleResponse{}
	err := client.Request("ess", "ExecuteScalingRule", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ExecuteScalingRuleResponse struct {
	ScalingActivityId goaliyun.String
	RequestId         goaliyun.String
}
