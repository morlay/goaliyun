package nas

import (
	"github.com/morlay/goaliyun"
)

type ModifyAccessRuleRequest struct {
	RWAccessType    string `position:"Query" name:"RWAccessType"`
	SourceCidrIp    string `position:"Query" name:"SourceCidrIp"`
	UserAccessType  string `position:"Query" name:"UserAccessType"`
	Priority        int64  `position:"Query" name:"Priority"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyAccessRuleRequest) Invoke(client goaliyun.Client) (*ModifyAccessRuleResponse, error) {
	resp := &ModifyAccessRuleResponse{}
	err := client.Request("nas", "ModifyAccessRule", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAccessRuleResponse struct {
	RequestId goaliyun.String
}
