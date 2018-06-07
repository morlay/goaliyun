package ram

import (
	"github.com/morlay/goaliyun"
)

type AttachPolicyToRoleRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	RoleName   string `position:"Query" name:"RoleName"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *AttachPolicyToRoleRequest) Invoke(client goaliyun.Client) (*AttachPolicyToRoleResponse, error) {
	resp := &AttachPolicyToRoleResponse{}
	err := client.Request("ram", "AttachPolicyToRole", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachPolicyToRoleResponse struct {
	RequestId goaliyun.String
}
