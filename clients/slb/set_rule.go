package slb

import (
	"github.com/morlay/goaliyun"
)

type SetRuleRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RuleId               string `position:"Query" name:"RuleId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetRuleRequest) Invoke(client goaliyun.Client) (*SetRuleResponse, error) {
	resp := &SetRuleResponse{}
	err := client.Request("slb", "SetRule", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetRuleResponse struct {
	RequestId goaliyun.String
}
