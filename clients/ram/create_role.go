package ram

import (
	"github.com/morlay/goaliyun"
)

type CreateRoleRequest struct {
	RoleName                 string `position:"Query" name:"RoleName"`
	Description              string `position:"Query" name:"Description"`
	AssumeRolePolicyDocument string `position:"Query" name:"AssumeRolePolicyDocument"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *CreateRoleRequest) Invoke(client goaliyun.Client) (*CreateRoleResponse, error) {
	resp := &CreateRoleResponse{}
	err := client.Request("ram", "CreateRole", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateRoleResponse struct {
	RequestId goaliyun.String
	Role      CreateRoleRole
}

type CreateRoleRole struct {
	RoleId                   goaliyun.String
	RoleName                 goaliyun.String
	Arn                      goaliyun.String
	Description              goaliyun.String
	AssumeRolePolicyDocument goaliyun.String
	CreateDate               goaliyun.String
}
