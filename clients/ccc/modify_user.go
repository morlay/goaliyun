package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyUserRequest struct {
	SkillLevels   *ModifyUserSkillLevelList   `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId    string                      `position:"Query" name:"InstanceId"`
	Phone         string                      `position:"Query" name:"Phone"`
	RoleIds       *ModifyUserRoleIdList       `position:"Query" type:"Repeated" name:"RoleId"`
	DisplayName   string                      `position:"Query" name:"DisplayName"`
	SkillGroupIds *ModifyUserSkillGroupIdList `position:"Query" type:"Repeated" name:"SkillGroupId"`
	UserId        string                      `position:"Query" name:"UserId"`
	Email         string                      `position:"Query" name:"Email"`
	RegionId      string                      `position:"Query" name:"RegionId"`
}

func (req *ModifyUserRequest) Invoke(client goaliyun.Client) (*ModifyUserResponse, error) {
	resp := &ModifyUserResponse{}
	err := client.Request("ccc", "ModifyUser", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyUserResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}

type ModifyUserSkillLevelList []int64

func (list *ModifyUserSkillLevelList) UnmarshalJSON(data []byte) error {
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

type ModifyUserRoleIdList []string

func (list *ModifyUserRoleIdList) UnmarshalJSON(data []byte) error {
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

type ModifyUserSkillGroupIdList []string

func (list *ModifyUserSkillGroupIdList) UnmarshalJSON(data []byte) error {
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
