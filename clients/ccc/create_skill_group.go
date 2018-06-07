package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateSkillGroupRequest struct {
	SkillLevels            *CreateSkillGroupSkillLevelList            `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId             string                                     `position:"Query" name:"InstanceId"`
	OutboundPhoneNumberIds *CreateSkillGroupOutboundPhoneNumberIdList `position:"Query" type:"Repeated" name:"OutboundPhoneNumberId"`
	Name                   string                                     `position:"Query" name:"Name"`
	Description            string                                     `position:"Query" name:"Description"`
	UserIds                *CreateSkillGroupUserIdList                `position:"Query" type:"Repeated" name:"UserId"`
	RegionId               string                                     `position:"Query" name:"RegionId"`
}

func (req *CreateSkillGroupRequest) Invoke(client goaliyun.Client) (*CreateSkillGroupResponse, error) {
	resp := &CreateSkillGroupResponse{}
	err := client.Request("ccc", "CreateSkillGroup", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSkillGroupResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	SkillGroupId   goaliyun.String
}

type CreateSkillGroupSkillLevelList []int64

func (list *CreateSkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
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

type CreateSkillGroupOutboundPhoneNumberIdList []string

func (list *CreateSkillGroupOutboundPhoneNumberIdList) UnmarshalJSON(data []byte) error {
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

type CreateSkillGroupUserIdList []string

func (list *CreateSkillGroupUserIdList) UnmarshalJSON(data []byte) error {
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
