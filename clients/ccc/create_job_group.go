package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateJobGroupRequest struct {
	CallingNumbers *CreateJobGroupCallingNumberList `position:"Query" type:"Repeated" name:"CallingNumber"`
	InstanceId     string                           `position:"Query" name:"InstanceId"`
	StrategyJson   string                           `position:"Query" name:"StrategyJson"`
	Name           string                           `position:"Query" name:"Name"`
	Description    string                           `position:"Query" name:"Description"`
	ScenarioId     string                           `position:"Query" name:"ScenarioId"`
	RegionId       string                           `position:"Query" name:"RegionId"`
}

func (req *CreateJobGroupRequest) Invoke(client goaliyun.Client) (*CreateJobGroupResponse, error) {
	resp := &CreateJobGroupResponse{}
	err := client.Request("ccc", "CreateJobGroup", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateJobGroupResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	JobGroup       CreateJobGroupJobGroup
}

type CreateJobGroupJobGroup struct {
	JobGroupId          goaliyun.String
	JobGroupName        goaliyun.String
	JobGroupDescription goaliyun.String
	ScenarioId          goaliyun.String
	JobFilePath         goaliyun.String
	CreationTime        goaliyun.Integer
	CallingNumbers      CreateJobGroupCallingNumberList
	Strategy            CreateJobGroupStrategy
}

type CreateJobGroupStrategy struct {
	StrategyId          goaliyun.String
	StrategyName        goaliyun.String
	StrategyDescription goaliyun.String
	Type                goaliyun.String
	StartTime           goaliyun.Integer
	EndTime             goaliyun.Integer
	RepeatBy            goaliyun.String
	MaxAttemptsPerDay   goaliyun.Integer
	MinAttemptInterval  goaliyun.Integer
	Customized          goaliyun.String
	RoutingStrategy     goaliyun.String
	FollowUpStrategy    goaliyun.String
	IsTemplate          bool
	WorkingTime         CreateJobGroupTimeFrameList
	RepeatDays          CreateJobGroupRepeatDayList
}

type CreateJobGroupTimeFrame struct {
	From goaliyun.String
	To   goaliyun.String
}

type CreateJobGroupCallingNumberList []goaliyun.String

func (list *CreateJobGroupCallingNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type CreateJobGroupTimeFrameList []CreateJobGroupTimeFrame

func (list *CreateJobGroupTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateJobGroupTimeFrame)
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

type CreateJobGroupRepeatDayList []goaliyun.String

func (list *CreateJobGroupRepeatDayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
