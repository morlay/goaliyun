package emr

import (
	"github.com/morlay/goaliyun"
)

type CheckUserRoleRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ConsoleRoleName string `position:"Query" name:"ConsoleRoleName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CheckUserRoleRequest) Invoke(client goaliyun.Client) (*CheckUserRoleResponse, error) {
	resp := &CheckUserRoleResponse{}
	err := client.Request("emr", "CheckUserRole", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckUserRoleResponse struct {
	RequestId goaliyun.String
	Success   bool
}
