package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteRulesRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RuleIds              string `position:"Query" name:"RuleIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteRulesRequest) Invoke(client goaliyun.Client) (*DeleteRulesResponse, error) {
	resp := &DeleteRulesResponse{}
	err := client.Request("slb", "DeleteRules", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteRulesResponse struct {
	RequestId goaliyun.String
}
