package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateRulesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int64  `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RuleList             string `position:"Query" name:"RuleList"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateRulesRequest) Invoke(client goaliyun.Client) (*CreateRulesResponse, error) {
	resp := &CreateRulesResponse{}
	err := client.Request("slb", "CreateRules", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateRulesResponse struct {
	RequestId goaliyun.String
	Rules     CreateRulesRuleList
}

type CreateRulesRule struct {
	RuleId   goaliyun.String
	RuleName goaliyun.String
}

type CreateRulesRuleList []CreateRulesRule

func (list *CreateRulesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateRulesRule)
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
