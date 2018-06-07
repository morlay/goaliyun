package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRolesRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListRolesRequest) Invoke(client goaliyun.Client) (*ListRolesResponse, error) {
	resp := &ListRolesResponse{}
	err := client.Request("ccc", "ListRoles", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRolesResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Roles          ListRolesRoleList
}

type ListRolesRole struct {
	RoleId          goaliyun.String
	InstanceId      goaliyun.String
	RoleName        goaliyun.String
	RoleDescription goaliyun.String
}

type ListRolesRoleList []ListRolesRole

func (list *ListRolesRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRolesRole)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
