package nas

import (
	"github.com/morlay/goaliyun"
)

type DeleteAccessRuleRequest struct {
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string `position:"Query" name:"AccessRuleId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteAccessRuleRequest) Invoke(client goaliyun.Client) (*DeleteAccessRuleResponse, error) {
	resp := &DeleteAccessRuleResponse{}
	err := client.Request("nas", "DeleteAccessRule", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAccessRuleResponse struct {
	RequestId goaliyun.String
}
