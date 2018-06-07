package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeScalingRulesRequest struct {
	ScalingRuleName1     string `position:"Query" name:"ScalingRuleName.1"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingRuleName2     string `position:"Query" name:"ScalingRuleName.2"`
	ScalingRuleName3     string `position:"Query" name:"ScalingRuleName.3"`
	ScalingRuleName4     string `position:"Query" name:"ScalingRuleName.4"`
	ScalingRuleName5     string `position:"Query" name:"ScalingRuleName.5"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	ScalingRuleName6     string `position:"Query" name:"ScalingRuleName.6"`
	ScalingRuleName7     string `position:"Query" name:"ScalingRuleName.7"`
	ScalingRuleName8     string `position:"Query" name:"ScalingRuleName.8"`
	ScalingRuleAri9      string `position:"Query" name:"ScalingRuleAri.9"`
	ScalingRuleName9     string `position:"Query" name:"ScalingRuleName.9"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ScalingRuleId10      string `position:"Query" name:"ScalingRuleId.10"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingRuleAri1      string `position:"Query" name:"ScalingRuleAri.1"`
	ScalingRuleAri2      string `position:"Query" name:"ScalingRuleAri.2"`
	ScalingRuleName10    string `position:"Query" name:"ScalingRuleName.10"`
	ScalingRuleAri3      string `position:"Query" name:"ScalingRuleAri.3"`
	ScalingRuleAri4      string `position:"Query" name:"ScalingRuleAri.4"`
	ScalingRuleId8       string `position:"Query" name:"ScalingRuleId.8"`
	ScalingRuleAri5      string `position:"Query" name:"ScalingRuleAri.5"`
	ScalingRuleId9       string `position:"Query" name:"ScalingRuleId.9"`
	ScalingRuleAri6      string `position:"Query" name:"ScalingRuleAri.6"`
	ScalingRuleAri7      string `position:"Query" name:"ScalingRuleAri.7"`
	ScalingRuleAri10     string `position:"Query" name:"ScalingRuleAri.10"`
	ScalingRuleAri8      string `position:"Query" name:"ScalingRuleAri.8"`
	ScalingRuleId4       string `position:"Query" name:"ScalingRuleId.4"`
	ScalingRuleId5       string `position:"Query" name:"ScalingRuleId.5"`
	ScalingRuleId6       string `position:"Query" name:"ScalingRuleId.6"`
	ScalingRuleId7       string `position:"Query" name:"ScalingRuleId.7"`
	ScalingRuleId1       string `position:"Query" name:"ScalingRuleId.1"`
	ScalingRuleId2       string `position:"Query" name:"ScalingRuleId.2"`
	ScalingRuleId3       string `position:"Query" name:"ScalingRuleId.3"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeScalingRulesRequest) Invoke(client goaliyun.Client) (*DescribeScalingRulesResponse, error) {
	resp := &DescribeScalingRulesResponse{}
	err := client.Request("ess", "DescribeScalingRules", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeScalingRulesResponse struct {
	TotalCount   goaliyun.Integer
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	RequestId    goaliyun.String
	ScalingRules DescribeScalingRulesScalingRuleList
}

type DescribeScalingRulesScalingRule struct {
	ScalingRuleId   goaliyun.String
	ScalingGroupId  goaliyun.String
	ScalingRuleName goaliyun.String
	Cooldown        goaliyun.Integer
	AdjustmentType  goaliyun.String
	AdjustmentValue goaliyun.Integer
	MinSize         goaliyun.Integer
	MaxSize         goaliyun.Integer
	ScalingRuleAri  goaliyun.String
}

type DescribeScalingRulesScalingRuleList []DescribeScalingRulesScalingRule

func (list *DescribeScalingRulesScalingRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingRulesScalingRule)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
