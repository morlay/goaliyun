package ram

import (
	"github.com/morlay/goaliyun"
)

type GetRoleRequest struct {
	RoleName string `position:"Query" name:"RoleName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetRoleRequest) Invoke(client goaliyun.Client) (*GetRoleResponse, error) {
	resp := &GetRoleResponse{}
	err := client.Request("ram", "GetRole", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetRoleResponse struct {
	RequestId goaliyun.String
	Role      GetRoleRole
}

type GetRoleRole struct {
	RoleId                   goaliyun.String
	RoleName                 goaliyun.String
	Arn                      goaliyun.String
	Description              goaliyun.String
	AssumeRolePolicyDocument goaliyun.String
	CreateDate               goaliyun.String
	UpdateDate               goaliyun.String
}
