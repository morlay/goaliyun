package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetJobGroupRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetJobGroupRequest) Invoke(client goaliyun.Client) (*GetJobGroupResponse, error) {
	resp := &GetJobGroupResponse{}
	err := client.Request("ccc", "GetJobGroup", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetJobGroupResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	JobGroup       GetJobGroupJobGroup
}

type GetJobGroupJobGroup struct {
	JobGroupId          goaliyun.String
	JobGroupName        goaliyun.String
	JobGroupDescription goaliyun.String
	ScenarioId          goaliyun.String
	JobFilePath         goaliyun.String
	CreationTime        goaliyun.Integer
	CallingNumbers      GetJobGroupCallingNumberList
	Strategy            GetJobGroupStrategy
	Progress            GetJobGroupProgress
}

type GetJobGroupStrategy struct {
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
	WorkingTime         GetJobGroupTimeFrameList
	RepeatDays          GetJobGroupRepeatDayList
}

type GetJobGroupTimeFrame struct {
	BeginTime goaliyun.String
	EndTime   goaliyun.String
}

type GetJobGroupProgress struct {
	TotalJobs        goaliyun.Integer
	Status           goaliyun.String
	TotalNotAnswered goaliyun.Integer
	TotalCompleted   goaliyun.Integer
	StartTime        goaliyun.Integer
	Duration         goaliyun.Integer
	Categories       GetJobGroupKeyValuePairList
}

type GetJobGroupKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type GetJobGroupCallingNumberList []goaliyun.String

func (list *GetJobGroupCallingNumberList) UnmarshalJSON(data []byte) error {
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

type GetJobGroupTimeFrameList []GetJobGroupTimeFrame

func (list *GetJobGroupTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobGroupTimeFrame)
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

type GetJobGroupRepeatDayList []goaliyun.String

func (list *GetJobGroupRepeatDayList) UnmarshalJSON(data []byte) error {
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

type GetJobGroupKeyValuePairList []GetJobGroupKeyValuePair

func (list *GetJobGroupKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobGroupKeyValuePair)
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
