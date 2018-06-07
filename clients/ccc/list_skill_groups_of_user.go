package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListSkillGroupsOfUserRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListSkillGroupsOfUserRequest) Invoke(client goaliyun.Client) (*ListSkillGroupsOfUserResponse, error) {
	resp := &ListSkillGroupsOfUserResponse{}
	err := client.Request("ccc", "ListSkillGroupsOfUser", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListSkillGroupsOfUserResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	SkillLevels    ListSkillGroupsOfUserSkillLevelList
}

type ListSkillGroupsOfUserSkillLevel struct {
	SkillLevelId goaliyun.String
	Level        goaliyun.Integer
	Skill        ListSkillGroupsOfUserSkill
}

type ListSkillGroupsOfUserSkill struct {
	SkillGroupId          goaliyun.String
	InstanceId            goaliyun.String
	SkillGroupName        goaliyun.String
	SkillGroupDescription goaliyun.String
	OutboundPhoneNumbers  ListSkillGroupsOfUserPhoneNumberList
}

type ListSkillGroupsOfUserPhoneNumber struct {
	PhoneNumberId          goaliyun.String
	InstanceId             goaliyun.String
	Number                 goaliyun.String
	PhoneNumberDescription goaliyun.String
	TestOnly               bool
	RemainingTime          goaliyun.Integer
	AllowOutbound          bool
	Usage                  goaliyun.String
	Trunks                 goaliyun.Integer
	Province               goaliyun.String
	City                   goaliyun.String
}

type ListSkillGroupsOfUserSkillLevelList []ListSkillGroupsOfUserSkillLevel

func (list *ListSkillGroupsOfUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsOfUserSkillLevel)
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

type ListSkillGroupsOfUserPhoneNumberList []ListSkillGroupsOfUserPhoneNumber

func (list *ListSkillGroupsOfUserPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsOfUserPhoneNumber)
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
