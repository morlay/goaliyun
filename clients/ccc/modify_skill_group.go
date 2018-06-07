package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifySkillGroupRequest struct {
	SkillLevels            *ModifySkillGroupSkillLevelList            `position:"Query" type:"Repeated" name:"SkillLevel"`
	InstanceId             string                                     `position:"Query" name:"InstanceId"`
	OutboundPhoneNumberIds *ModifySkillGroupOutboundPhoneNumberIdList `position:"Query" type:"Repeated" name:"OutboundPhoneNumberId"`
	SkillGroupId           string                                     `position:"Query" name:"SkillGroupId"`
	Name                   string                                     `position:"Query" name:"Name"`
	Description            string                                     `position:"Query" name:"Description"`
	UserIds                *ModifySkillGroupUserIdList                `position:"Query" type:"Repeated" name:"UserId"`
	RegionId               string                                     `position:"Query" name:"RegionId"`
}

func (req *ModifySkillGroupRequest) Invoke(client goaliyun.Client) (*ModifySkillGroupResponse, error) {
	resp := &ModifySkillGroupResponse{}
	err := client.Request("ccc", "ModifySkillGroup", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySkillGroupResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}

type ModifySkillGroupSkillLevelList []int64

func (list *ModifySkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
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

type ModifySkillGroupOutboundPhoneNumberIdList []string

func (list *ModifySkillGroupOutboundPhoneNumberIdList) UnmarshalJSON(data []byte) error {
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

type ModifySkillGroupUserIdList []string

func (list *ModifySkillGroupUserIdList) UnmarshalJSON(data []byte) error {
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
