package emr

import (
	"github.com/morlay/goaliyun"
)

type PassRoleRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ConsoleRoleName string `position:"Query" name:"ConsoleRoleName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *PassRoleRequest) Invoke(client goaliyun.Client) (*PassRoleResponse, error) {
	resp := &PassRoleResponse{}
	err := client.Request("emr", "PassRole", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PassRoleResponse struct {
	RequestId goaliyun.String
	Success   bool
}
