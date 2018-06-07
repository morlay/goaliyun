package ram

import (
	"github.com/morlay/goaliyun"
)

type DeleteRoleRequest struct {
	RoleName string `position:"Query" name:"RoleName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteRoleRequest) Invoke(client goaliyun.Client) (*DeleteRoleResponse, error) {
	resp := &DeleteRoleResponse{}
	err := client.Request("ram", "DeleteRole", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteRoleResponse struct {
	RequestId goaliyun.String
}
