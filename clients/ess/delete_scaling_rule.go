package ess

import (
	"github.com/morlay/goaliyun"
)

type DeleteScalingRuleRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingRuleId        string `position:"Query" name:"ScalingRuleId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteScalingRuleRequest) Invoke(client goaliyun.Client) (*DeleteScalingRuleResponse, error) {
	resp := &DeleteScalingRuleResponse{}
	err := client.Request("ess", "DeleteScalingRule", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteScalingRuleResponse struct {
	RequestId goaliyun.String
}
