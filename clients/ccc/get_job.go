package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetJobRequest struct {
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetJobRequest) Invoke(client goaliyun.Client) (*GetJobResponse, error) {
	resp := &GetJobResponse{}
	err := client.Request("ccc", "GetJob", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetJobResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Job            GetJobJob
}

type GetJobJob struct {
	JobId          goaliyun.String
	GroupId        goaliyun.String
	ScenarioId     goaliyun.String
	StrategyId     goaliyun.String
	Priority       goaliyun.Integer
	SystemPriority goaliyun.Integer
	Status         goaliyun.String
	ReferenceId    goaliyun.String
	FailureReason  goaliyun.String
	Contacts       GetJobContactList
	Extras         GetJobKeyValuePairList
	Tasks          GetJobTaskList
	Summary        GetJobSummaryItem3List
	CallingNumbers GetJobCallingNumberList
}

type GetJobContact struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
}

type GetJobKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type GetJobTask struct {
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
	Conversation  GetJobConversationDetailList
	Contact2      GetJobContact2
}

type GetJobConversationDetail struct {
	Timestamp goaliyun.Integer
	Speaker   goaliyun.String
	Script    goaliyun.String
	Summary1  GetJobSummaryItemList
}

type GetJobSummaryItem struct {
	Category    goaliyun.String
	SummaryName goaliyun.String
	Content     goaliyun.String
}

type GetJobContact2 struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
}

type GetJobSummaryItem3 struct {
	Category    goaliyun.String
	SummaryName goaliyun.String
	Content     goaliyun.String
}

type GetJobContactList []GetJobContact

func (list *GetJobContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobContact)
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

type GetJobKeyValuePairList []GetJobKeyValuePair

func (list *GetJobKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobKeyValuePair)
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

type GetJobTaskList []GetJobTask

func (list *GetJobTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobTask)
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

type GetJobSummaryItem3List []GetJobSummaryItem3

func (list *GetJobSummaryItem3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobSummaryItem3)
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

type GetJobCallingNumberList []goaliyun.String

func (list *GetJobCallingNumberList) UnmarshalJSON(data []byte) error {
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

type GetJobConversationDetailList []GetJobConversationDetail

func (list *GetJobConversationDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobConversationDetail)
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

type GetJobSummaryItemList []GetJobSummaryItem

func (list *GetJobSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobSummaryItem)
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
