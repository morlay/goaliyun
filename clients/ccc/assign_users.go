package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AssignUsersRequest struct {
	UserRamIds    *AssignUsersUserRamIdList    `position:"Query" type:"Repeated" name:"UserRamId"`
	SkillLevels   *AssignUsersSkillLevelList   `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId    string                       `position:"Query" name:"InstanceId"`
	RoleIds       *AssignUsersRoleIdList       `position:"Query" type:"Repeated" name:"RoleId"`
	SkillGroupIds *AssignUsersSkillGroupIdList `position:"Query" type:"Repeated" name:"SkillGroupId"`
	RegionId      string                       `position:"Query" name:"RegionId"`
}

func (req *AssignUsersRequest) Invoke(client goaliyun.Client) (*AssignUsersResponse, error) {
	resp := &AssignUsersResponse{}
	err := client.Request("ccc", "AssignUsers", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssignUsersResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}

type AssignUsersUserRamIdList []string

func (list *AssignUsersUserRamIdList) UnmarshalJSON(data []byte) error {
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

type AssignUsersSkillLevelList []int64

func (list *AssignUsersSkillLevelList) UnmarshalJSON(data []byte) error {
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

type AssignUsersRoleIdList []string

func (list *AssignUsersRoleIdList) UnmarshalJSON(data []byte) error {
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

type AssignUsersSkillGroupIdList []string

func (list *AssignUsersSkillGroupIdList) UnmarshalJSON(data []byte) error {
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
