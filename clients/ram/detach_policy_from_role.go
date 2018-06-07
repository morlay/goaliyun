package ram

import (
	"github.com/morlay/goaliyun"
)

type DetachPolicyFromRoleRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	RoleName   string `position:"Query" name:"RoleName"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DetachPolicyFromRoleRequest) Invoke(client goaliyun.Client) (*DetachPolicyFromRoleResponse, error) {
	resp := &DetachPolicyFromRoleResponse{}
	err := client.Request("ram", "DetachPolicyFromRole", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachPolicyFromRoleResponse struct {
	RequestId goaliyun.String
}
