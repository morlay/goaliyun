package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetUserRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetUserRequest) Invoke(client goaliyun.Client) (*GetUserResponse, error) {
	resp := &GetUserResponse{}
	err := client.Request("ccc", "GetUser", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	User           GetUserUser
}

type GetUserUser struct {
	UserId      goaliyun.String
	RamId       goaliyun.String
	InstanceId  goaliyun.String
	Roles       GetUserRoleList
	SkillLevels GetUserSkillLevelList
	Detail      GetUserDetail
}

type GetUserRole struct {
	RoleId          goaliyun.String
	InstanceId      goaliyun.String
	RoleName        goaliyun.String
	RoleDescription goaliyun.String
}

type GetUserSkillLevel struct {
	SkillLevelId goaliyun.String
	Level        goaliyun.Integer
	Skill        GetUserSkill
}

type GetUserSkill struct {
	SkillGroupId          goaliyun.String
	InstanceId            goaliyun.String
	SkillGroupName        goaliyun.String
	SkillGroupDescription goaliyun.String
}

type GetUserDetail struct {
	LoginName   goaliyun.String
	DisplayName goaliyun.String
	Phone       goaliyun.String
	Email       goaliyun.String
	Department  goaliyun.String
}

type GetUserRoleList []GetUserRole

func (list *GetUserRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserRole)
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

type GetUserSkillLevelList []GetUserSkillLevel

func (list *GetUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserSkillLevel)
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
