package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListSkillGroupsRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListSkillGroupsRequest) Invoke(client goaliyun.Client) (*ListSkillGroupsResponse, error) {
	resp := &ListSkillGroupsResponse{}
	err := client.Request("ccc", "ListSkillGroups", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListSkillGroupsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	SkillGroups    ListSkillGroupsSkillGroupList
}

type ListSkillGroupsSkillGroup struct {
	SkillGroupId          goaliyun.String
	InstanceId            goaliyun.String
	SkillGroupName        goaliyun.String
	AccSkillGroupName     goaliyun.String
	AccQueueName          goaliyun.String
	SkillGroupDescription goaliyun.String
	UserCount             goaliyun.Integer
	OutboundPhoneNumbers  ListSkillGroupsPhoneNumberList
}

type ListSkillGroupsPhoneNumber struct {
	PhoneNumberId          goaliyun.String
	InstanceId             goaliyun.String
	Number                 goaliyun.String
	PhoneNumberDescription goaliyun.String
	TestOnly               bool
	RemainingTime          goaliyun.Integer
	AllowOutbound          bool
	Usage                  goaliyun.String
	Trunks                 goaliyun.Integer
}

type ListSkillGroupsSkillGroupList []ListSkillGroupsSkillGroup

func (list *ListSkillGroupsSkillGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsSkillGroup)
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

type ListSkillGroupsPhoneNumberList []ListSkillGroupsPhoneNumber

func (list *ListSkillGroupsPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsPhoneNumber)
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
