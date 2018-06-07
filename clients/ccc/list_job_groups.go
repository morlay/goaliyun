package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobGroupsRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListJobGroupsRequest) Invoke(client goaliyun.Client) (*ListJobGroupsResponse, error) {
	resp := &ListJobGroupsResponse{}
	err := client.Request("ccc", "ListJobGroups", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobGroupsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	JobGroups      ListJobGroupsJobGroups
}

type ListJobGroupsJobGroups struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       ListJobGroupsJobGroupList
}

type ListJobGroupsJobGroup struct {
	JobGroupId          goaliyun.String
	JobGroupName        goaliyun.String
	JobGroupDescription goaliyun.String
	ScenarioId          goaliyun.String
	JobFilePath         goaliyun.String
	CreationTime        goaliyun.Integer
	CallingNumbers      ListJobGroupsCallingNumberList
	Strategy            ListJobGroupsStrategy
	Progress            ListJobGroupsProgress
}

type ListJobGroupsStrategy struct {
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
	WorkingTime         ListJobGroupsTimeFrameList
	RepeatDays          ListJobGroupsRepeatDayList
}

type ListJobGroupsTimeFrame struct {
	From goaliyun.String
	To   goaliyun.String
}

type ListJobGroupsProgress struct {
	TotalJobs        goaliyun.Integer
	Status           goaliyun.String
	TotalNotAnswered goaliyun.Integer
	TotalCompleted   goaliyun.Integer
	StartTime        goaliyun.Integer
	Duration         goaliyun.Integer
	Categories       ListJobGroupsKeyValuePairList
}

type ListJobGroupsKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type ListJobGroupsJobGroupList []ListJobGroupsJobGroup

func (list *ListJobGroupsJobGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobGroupsJobGroup)
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

type ListJobGroupsCallingNumberList []goaliyun.String

func (list *ListJobGroupsCallingNumberList) UnmarshalJSON(data []byte) error {
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

type ListJobGroupsTimeFrameList []ListJobGroupsTimeFrame

func (list *ListJobGroupsTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobGroupsTimeFrame)
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

type ListJobGroupsRepeatDayList []goaliyun.String

func (list *ListJobGroupsRepeatDayList) UnmarshalJSON(data []byte) error {
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

type ListJobGroupsKeyValuePairList []ListJobGroupsKeyValuePair

func (list *ListJobGroupsKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobGroupsKeyValuePair)
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
