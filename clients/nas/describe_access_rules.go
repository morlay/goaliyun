package nas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccessRulesRequest struct {
	PageSize        int64  `position:"Query" name:"PageSize"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccessRulesRequest) Invoke(client goaliyun.Client) (*DescribeAccessRulesResponse, error) {
	resp := &DescribeAccessRulesResponse{}
	err := client.Request("nas", "DescribeAccessRules", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccessRulesResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageSize    goaliyun.Integer
	PageNumber  goaliyun.Integer
	AccessRules DescribeAccessRulesAccessRuleList
}

type DescribeAccessRulesAccessRule struct {
	SourceCidrIp goaliyun.String
	Priority     goaliyun.Integer
	AccessRuleId goaliyun.String
	RWAccess     goaliyun.String
	UserAccess   goaliyun.String
}

type DescribeAccessRulesAccessRuleList []DescribeAccessRulesAccessRule

func (list *DescribeAccessRulesAccessRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessRulesAccessRule)
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
