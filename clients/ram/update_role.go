package ram

import (
	"github.com/morlay/goaliyun"
)

type UpdateRoleRequest struct {
	NewAssumeRolePolicyDocument string `position:"Query" name:"NewAssumeRolePolicyDocument"`
	RoleName                    string `position:"Query" name:"RoleName"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *UpdateRoleRequest) Invoke(client goaliyun.Client) (*UpdateRoleResponse, error) {
	resp := &UpdateRoleResponse{}
	err := client.Request("ram", "UpdateRole", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateRoleResponse struct {
	RequestId goaliyun.String
	Role      UpdateRoleRole
}

type UpdateRoleRole struct {
	RoleId                   goaliyun.String
	RoleName                 goaliyun.String
	Arn                      goaliyun.String
	Description              goaliyun.String
	AssumeRolePolicyDocument goaliyun.String
	CreateDate               goaliyun.String
	UpdateDate               goaliyun.String
}
