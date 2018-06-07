package ess

import (
	"github.com/morlay/goaliyun"
)

type ModifyScalingRuleRequest struct {
	ScalingRuleName      string `position:"Query" name:"ScalingRuleName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentValue      int64  `position:"Query" name:"AdjustmentValue"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Cooldown             int64  `position:"Query" name:"Cooldown"`
	AdjustmentType       string `position:"Query" name:"AdjustmentType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingRuleId        string `position:"Query" name:"ScalingRuleId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyScalingRuleRequest) Invoke(client goaliyun.Client) (*ModifyScalingRuleResponse, error) {
	resp := &ModifyScalingRuleResponse{}
	err := client.Request("ess", "ModifyScalingRule", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyScalingRuleResponse struct {
	RequestId goaliyun.String
}
