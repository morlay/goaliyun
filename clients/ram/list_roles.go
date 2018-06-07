package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRolesRequest struct {
	Marker   string `position:"Query" name:"Marker"`
	MaxItems int64  `position:"Query" name:"MaxItems"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListRolesRequest) Invoke(client goaliyun.Client) (*ListRolesResponse, error) {
	resp := &ListRolesResponse{}
	err := client.Request("ram", "ListRoles", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRolesResponse struct {
	RequestId   goaliyun.String
	IsTruncated bool
	Marker      goaliyun.String
	Roles       ListRolesRoleList
}

type ListRolesRole struct {
	RoleId      goaliyun.String
	RoleName    goaliyun.String
	Arn         goaliyun.String
	Description goaliyun.String
	CreateDate  goaliyun.String
	UpdateDate  goaliyun.String
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
