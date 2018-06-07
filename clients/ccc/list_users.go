package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListUsersRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListUsersRequest) Invoke(client goaliyun.Client) (*ListUsersResponse, error) {
	resp := &ListUsersResponse{}
	err := client.Request("ccc", "ListUsers", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListUsersResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Users          ListUsersUsers
}

type ListUsersUsers struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       ListUsersUserList
}

type ListUsersUser struct {
	UserId      goaliyun.String
	RamId       goaliyun.String
	InstanceId  goaliyun.String
	Primary     bool
	Roles       ListUsersRoleList
	SkillLevels ListUsersSkillLevelList
	Detail      ListUsersDetail
}

type ListUsersRole struct {
	RoleId          goaliyun.String
	InstanceId      goaliyun.String
	RoleName        goaliyun.String
	RoleDescription goaliyun.String
}

type ListUsersSkillLevel struct {
	SkillLevelId goaliyun.String
	Level        goaliyun.Integer
	Skill        ListUsersSkill
}

type ListUsersSkill struct {
	SkillGroupId          goaliyun.String
	InstanceId            goaliyun.String
	SkillGroupName        goaliyun.String
	SkillGroupDescription goaliyun.String
}

type ListUsersDetail struct {
	LoginName   goaliyun.String
	DisplayName goaliyun.String
	Phone       goaliyun.String
	Email       goaliyun.String
	Department  goaliyun.String
}

type ListUsersUserList []ListUsersUser

func (list *ListUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersUser)
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

type ListUsersRoleList []ListUsersRole

func (list *ListUsersRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersRole)
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

type ListUsersSkillLevelList []ListUsersSkillLevel

func (list *ListUsersSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersSkillLevel)
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
