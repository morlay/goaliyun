package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRealTimeAgentRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListRealTimeAgentRequest) Invoke(client goaliyun.Client) (*ListRealTimeAgentResponse, error) {
	resp := &ListRealTimeAgentResponse{}
	err := client.Request("ccc", "ListRealTimeAgent", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRealTimeAgentResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Data           ListRealTimeAgentUserList
}

type ListRealTimeAgentUser struct {
	RamId       goaliyun.String
	DisplayName goaliyun.String
	Phone       goaliyun.String
	Dn          goaliyun.String
	State       goaliyun.String
	StateDesc   goaliyun.String
	SkillLevels ListRealTimeAgentSkillLevelList
}

type ListRealTimeAgentSkillLevel struct {
	SkillLevelId goaliyun.String
	Level        goaliyun.Integer
	Skill        ListRealTimeAgentSkill
}

type ListRealTimeAgentSkill struct {
	SkillGroupId          goaliyun.String
	InstanceId            goaliyun.String
	SkillGroupName        goaliyun.String
	SkillGroupDescription goaliyun.String
}

type ListRealTimeAgentUserList []ListRealTimeAgentUser

func (list *ListRealTimeAgentUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRealTimeAgentUser)
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

type ListRealTimeAgentSkillLevelList []ListRealTimeAgentSkillLevel

func (list *ListRealTimeAgentSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRealTimeAgentSkillLevel)
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
