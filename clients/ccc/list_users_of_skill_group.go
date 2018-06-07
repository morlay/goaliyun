package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListUsersOfSkillGroupRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ListUsersOfSkillGroupRequest) Invoke(client goaliyun.Client) (*ListUsersOfSkillGroupResponse, error) {
	resp := &ListUsersOfSkillGroupResponse{}
	err := client.Request("ccc", "ListUsersOfSkillGroup", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListUsersOfSkillGroupResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Users          ListUsersOfSkillGroupUsers
}

type ListUsersOfSkillGroupUsers struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       ListUsersOfSkillGroupUserList
}

type ListUsersOfSkillGroupUser struct {
	UserId      goaliyun.String
	RamId       goaliyun.String
	InstanceId  goaliyun.String
	Roles       ListUsersOfSkillGroupRoleList
	SkillLevels ListUsersOfSkillGroupSkillLevelList
	Detail      ListUsersOfSkillGroupDetail
}

type ListUsersOfSkillGroupRole struct {
	RoleId          goaliyun.String
	InstanceId      goaliyun.String
	RoleName        goaliyun.String
	RoleDescription goaliyun.String
	UserCount       goaliyun.Integer
	Privileges      ListUsersOfSkillGroupPrivilegeList
}

type ListUsersOfSkillGroupPrivilege struct {
	PrivilegeId          goaliyun.String
	PrivilegeName        goaliyun.String
	PrivilegeDescription goaliyun.String
}

type ListUsersOfSkillGroupSkillLevel struct {
	SkillLevelId goaliyun.String
	Level        goaliyun.Integer
	Skill        ListUsersOfSkillGroupSkill
}

type ListUsersOfSkillGroupSkill struct {
	SkillGroupId          goaliyun.String
	InstanceId            goaliyun.String
	SkillGroupName        goaliyun.String
	SkillGroupDescription goaliyun.String
}

type ListUsersOfSkillGroupDetail struct {
	LoginName   goaliyun.String
	DisplayName goaliyun.String
	Phone       goaliyun.String
	Email       goaliyun.String
	Department  goaliyun.String
}

type ListUsersOfSkillGroupUserList []ListUsersOfSkillGroupUser

func (list *ListUsersOfSkillGroupUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupUser)
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

type ListUsersOfSkillGroupRoleList []ListUsersOfSkillGroupRole

func (list *ListUsersOfSkillGroupRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupRole)
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

type ListUsersOfSkillGroupSkillLevelList []ListUsersOfSkillGroupSkillLevel

func (list *ListUsersOfSkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupSkillLevel)
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

type ListUsersOfSkillGroupPrivilegeList []ListUsersOfSkillGroupPrivilege

func (list *ListUsersOfSkillGroupPrivilegeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupPrivilege)
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
