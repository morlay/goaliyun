package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetJobStatusByCallIdRequest struct {
	CallId     string `position:"Query" name:"CallId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetJobStatusByCallIdRequest) Invoke(client goaliyun.Client) (*GetJobStatusByCallIdResponse, error) {
	resp := &GetJobStatusByCallIdResponse{}
	err := client.Request("ccc", "GetJobStatusByCallId", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetJobStatusByCallIdResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Job            GetJobStatusByCallIdJob
}

type GetJobStatusByCallIdJob struct {
	JobId          goaliyun.String
	GroupId        goaliyun.String
	ScenarioId     goaliyun.String
	StrategyId     goaliyun.String
	Priority       goaliyun.Integer
	Status         goaliyun.String
	ReferenceId    goaliyun.String
	FailureReason  goaliyun.String
	Contacts       GetJobStatusByCallIdContactList
	Extras         GetJobStatusByCallIdKeyValuePairList
	Tasks          GetJobStatusByCallIdTaskList
	Summary        GetJobStatusByCallIdSummaryItem3List
	CallingNumbers GetJobStatusByCallIdCallingNumberList
}

type GetJobStatusByCallIdContact struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
	JobId       goaliyun.String
}

type GetJobStatusByCallIdKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type GetJobStatusByCallIdTask struct {
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
	Conversation  GetJobStatusByCallIdConversationDetailList
	Contact2      GetJobStatusByCallIdContact2
}

type GetJobStatusByCallIdConversationDetail struct {
	ConversationDetailId goaliyun.String
	TaskId               goaliyun.String
	Timestamp            goaliyun.Integer
	Speaker              goaliyun.String
	Script               goaliyun.String
	Summary1             GetJobStatusByCallIdSummaryItemList
}

type GetJobStatusByCallIdSummaryItem struct {
	SummaryId   goaliyun.String
	Category    goaliyun.String
	SummaryName goaliyun.String
	Content     goaliyun.String
}

type GetJobStatusByCallIdContact2 struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
	JobId       goaliyun.String
}

type GetJobStatusByCallIdSummaryItem3 struct {
	SummaryId   goaliyun.String
	Category    goaliyun.String
	SummaryName goaliyun.String
	Content     goaliyun.String
}

type GetJobStatusByCallIdContactList []GetJobStatusByCallIdContact

func (list *GetJobStatusByCallIdContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdContact)
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

type GetJobStatusByCallIdKeyValuePairList []GetJobStatusByCallIdKeyValuePair

func (list *GetJobStatusByCallIdKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdKeyValuePair)
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

type GetJobStatusByCallIdTaskList []GetJobStatusByCallIdTask

func (list *GetJobStatusByCallIdTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdTask)
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

type GetJobStatusByCallIdSummaryItem3List []GetJobStatusByCallIdSummaryItem3

func (list *GetJobStatusByCallIdSummaryItem3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdSummaryItem3)
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

type GetJobStatusByCallIdCallingNumberList []goaliyun.String

func (list *GetJobStatusByCallIdCallingNumberList) UnmarshalJSON(data []byte) error {
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

type GetJobStatusByCallIdConversationDetailList []GetJobStatusByCallIdConversationDetail

func (list *GetJobStatusByCallIdConversationDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdConversationDetail)
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

type GetJobStatusByCallIdSummaryItemList []GetJobStatusByCallIdSummaryItem

func (list *GetJobStatusByCallIdSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdSummaryItem)
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
