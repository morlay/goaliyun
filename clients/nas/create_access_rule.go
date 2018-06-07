package nas

import (
	"github.com/morlay/goaliyun"
)

type CreateAccessRuleRequest struct {
	RWAccessType    string `position:"Query" name:"RWAccessType"`
	SourceCidrIp    string `position:"Query" name:"SourceCidrIp"`
	UserAccessType  string `position:"Query" name:"UserAccessType"`
	Priority        int64  `position:"Query" name:"Priority"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateAccessRuleRequest) Invoke(client goaliyun.Client) (*CreateAccessRuleResponse, error) {
	resp := &CreateAccessRuleResponse{}
	err := client.Request("nas", "CreateAccessRule", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAccessRuleResponse struct {
	RequestId    goaliyun.String
	AccessRuleId goaliyun.String
}
