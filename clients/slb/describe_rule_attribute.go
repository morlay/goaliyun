package slb

import (
	"github.com/morlay/goaliyun"
)

type DescribeRuleAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RuleId               string `position:"Query" name:"RuleId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRuleAttributeRequest) Invoke(client goaliyun.Client) (*DescribeRuleAttributeResponse, error) {
	resp := &DescribeRuleAttributeResponse{}
	err := client.Request("slb", "DescribeRuleAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRuleAttributeResponse struct {
	RequestId      goaliyun.String
	RuleName       goaliyun.String
	LoadBalancerId goaliyun.String
	ListenerPort   goaliyun.String
	Domain         goaliyun.String
	Url            goaliyun.String
	VServerGroupId goaliyun.String
}
