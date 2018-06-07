package ess

import (
	"github.com/morlay/goaliyun"
)

type CreateScalingRuleRequest struct {
	ScalingRuleName      string `position:"Query" name:"ScalingRuleName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentValue      int64  `position:"Query" name:"AdjustmentValue"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Cooldown             int64  `position:"Query" name:"Cooldown"`
	AdjustmentType       string `position:"Query" name:"AdjustmentType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateScalingRuleRequest) Invoke(client goaliyun.Client) (*CreateScalingRuleResponse, error) {
	resp := &CreateScalingRuleResponse{}
	err := client.Request("ess", "CreateScalingRule", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateScalingRuleResponse struct {
	ScalingRuleId  goaliyun.String
	ScalingRuleAri goaliyun.String
	RequestId      goaliyun.String
}
