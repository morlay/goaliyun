package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateUserRequest struct {
	SkillLevels   *CreateUserSkillLevelList   `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId    string                      `position:"Query" name:"InstanceId"`
	LoginName     string                      `position:"Query" name:"LoginName"`
	Phone         string                      `position:"Query" name:"Phone"`
	RoleIds       *CreateUserRoleIdList       `position:"Query" type:"Repeated" name:"RoleId"`
	DisplayName   string                      `position:"Query" name:"DisplayName"`
	SkillGroupIds *CreateUserSkillGroupIdList `position:"Query" type:"Repeated" name:"SkillGroupId"`
	Email         string                      `position:"Query" name:"Email"`
	RegionId      string                      `position:"Query" name:"RegionId"`
}

func (req *CreateUserRequest) Invoke(client goaliyun.Client) (*CreateUserResponse, error) {
	resp := &CreateUserResponse{}
	err := client.Request("ccc", "CreateUser", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateUserResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	UserId         goaliyun.String
}

type CreateUserSkillLevelList []int64

func (list *CreateUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type CreateUserRoleIdList []string

func (list *CreateUserRoleIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type CreateUserSkillGroupIdList []string

func (list *CreateUserSkillGroupIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
