package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobStatusRequest struct {
	ContactName   string `position:"Query" name:"ContactName"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	TimeAlignment string `position:"Query" name:"TimeAlignment"`
	GroupId       string `position:"Query" name:"GroupId"`
	PhoneNumber   string `position:"Query" name:"PhoneNumber"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	EndTime       int64  `position:"Query" name:"EndTime"`
	StartTime     int64  `position:"Query" name:"StartTime"`
	ScenarioId    string `position:"Query" name:"ScenarioId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ListJobStatusRequest) Invoke(client goaliyun.Client) (*ListJobStatusResponse, error) {
	resp := &ListJobStatusResponse{}
	err := client.Request("ccc", "ListJobStatus", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobStatusResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Jobs           ListJobStatusJobs
}

type ListJobStatusJobs struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       ListJobStatusJobList
}

type ListJobStatusJob struct {
	JobId          goaliyun.String
	GroupId        goaliyun.String
	ScenarioId     goaliyun.String
	StrategyId     goaliyun.String
	Priority       goaliyun.Integer
	Status         goaliyun.String
	ReferenceId    goaliyun.String
	FailureReason  goaliyun.String
	Contacts       ListJobStatusContactList
	Extras         ListJobStatusKeyValuePairList
	Tasks          ListJobStatusTaskList
	Summary        ListJobStatusSummaryItemList
	CallingNumbers ListJobStatusCallingNumberList
}

type ListJobStatusContact struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
	JobId       goaliyun.String
}

type ListJobStatusKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type ListJobStatusTask struct {
	TaskId        goaliyun.String
	JobId         goaliyun.String
	ScenarioId    goaliyun.String
	ChatbotId     goaliyun.String
	PlanedTime    goaliyun.Integer
	ActualTime    goaliyun.Integer
	CallingNumber goaliyun.String
	CalledNumber  goaliyun.String
	CallId        goaliyun.String
	Status        goaliyun.String
	Brief         goaliyun.String
	Duration      goaliyun.Integer
	Contact1      ListJobStatusContact1
}

type ListJobStatusContact1 struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
	JobId       goaliyun.String
}

type ListJobStatusSummaryItem struct {
	SummaryId            goaliyun.String
	GroupId              goaliyun.String
	JobId                goaliyun.String
	TaskId               goaliyun.String
	ConversationDetailId goaliyun.String
	Category             goaliyun.String
	SummaryName          goaliyun.String
	Content              goaliyun.String
}

type ListJobStatusJobList []ListJobStatusJob

func (list *ListJobStatusJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusJob)
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

type ListJobStatusContactList []ListJobStatusContact

func (list *ListJobStatusContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusContact)
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

type ListJobStatusKeyValuePairList []ListJobStatusKeyValuePair

func (list *ListJobStatusKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusKeyValuePair)
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

type ListJobStatusTaskList []ListJobStatusTask

func (list *ListJobStatusTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusTask)
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

type ListJobStatusSummaryItemList []ListJobStatusSummaryItem

func (list *ListJobStatusSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusSummaryItem)
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

type ListJobStatusCallingNumberList []goaliyun.String

func (list *ListJobStatusCallingNumberList) UnmarshalJSON(data []byte) error {
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
