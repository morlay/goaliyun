package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListHealthRuleResultRequest struct {
	Component       string `position:"Query" name:"Component"`
	HostName        string `position:"Query" name:"HostName"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Pass            int64  `position:"Query" name:"Pass"`
	Service         string `position:"Query" name:"Service"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListHealthRuleResultRequest) Invoke(client goaliyun.Client) (*ListHealthRuleResultResponse, error) {
	resp := &ListHealthRuleResultResponse{}
	err := client.Request("emr", "ListHealthRuleResult", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListHealthRuleResultResponse struct {
	RequestId            goaliyun.String
	Total                goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	HealthRuleResultList ListHealthRuleResultHealthRuleResultList
}

type ListHealthRuleResultHealthRuleResult struct {
	Id              goaliyun.Integer
	ClusterId       goaliyun.Integer
	RuleId          goaliyun.Integer
	RuleName        goaliyun.String
	RuleTitle       goaliyun.String
	RuleStatus      goaliyun.String
	RuleDescription goaliyun.String
	Service         goaliyun.String
	Component       goaliyun.String
	Pass            goaliyun.String
	HostNames       goaliyun.String
}

type ListHealthRuleResultHealthRuleResultList []ListHealthRuleResultHealthRuleResult

func (list *ListHealthRuleResultHealthRuleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListHealthRuleResultHealthRuleResult)
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
