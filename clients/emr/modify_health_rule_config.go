package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyHealthRuleConfigRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RuleName        string `position:"Query" name:"RuleName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RuleId          int64  `position:"Query" name:"RuleId"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyHealthRuleConfigRequest) Invoke(client goaliyun.Client) (*ModifyHealthRuleConfigResponse, error) {
	resp := &ModifyHealthRuleConfigResponse{}
	err := client.Request("emr", "ModifyHealthRuleConfig", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyHealthRuleConfigResponse struct {
	RequestId goaliyun.String
}
