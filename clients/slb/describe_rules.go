package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRulesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int64  `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRulesRequest) Invoke(client goaliyun.Client) (*DescribeRulesResponse, error) {
	resp := &DescribeRulesResponse{}
	err := client.Request("slb", "DescribeRules", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRulesResponse struct {
	RequestId goaliyun.String
	Rules     DescribeRulesRuleList
}

type DescribeRulesRule struct {
	RuleId         goaliyun.String
	RuleName       goaliyun.String
	Domain         goaliyun.String
	Url            goaliyun.String
	VServerGroupId goaliyun.String
}

type DescribeRulesRuleList []DescribeRulesRule

func (list *DescribeRulesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRulesRule)
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
