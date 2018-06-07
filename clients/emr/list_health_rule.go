package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListHealthRuleRequest struct {
	Component       string `position:"Query" name:"Component"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Service         string `position:"Query" name:"Service"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListHealthRuleRequest) Invoke(client goaliyun.Client) (*ListHealthRuleResponse, error) {
	resp := &ListHealthRuleResponse{}
	err := client.Request("emr", "ListHealthRule", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListHealthRuleResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Total      goaliyun.Integer
	Rule       ListHealthRuleRuleItemList
}

type ListHealthRuleRuleItem struct {
	Id          goaliyun.Integer
	Name        goaliyun.String
	Status      goaliyun.String
	Service     goaliyun.String
	Component   goaliyun.String
	Title       goaliyun.String
	Description goaliyun.String
	Explanation goaliyun.String
	Solution    goaliyun.String
}

type ListHealthRuleRuleItemList []ListHealthRuleRuleItem

func (list *ListHealthRuleRuleItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListHealthRuleRuleItem)
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
