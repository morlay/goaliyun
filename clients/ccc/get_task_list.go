package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetTaskListRequest struct {
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetTaskListRequest) Invoke(client goaliyun.Client) (*GetTaskListResponse, error) {
	resp := &GetTaskListResponse{}
	err := client.Request("ccc", "GetTaskList", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetTaskListResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Tasks          GetTaskListTaskList
}

type GetTaskListTask struct {
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
	Conversation  GetTaskListConversationDetailList
	Contact       GetTaskListContact
}

type GetTaskListConversationDetail struct {
	Timestamp goaliyun.Integer
	Speaker   goaliyun.String
	Script    goaliyun.String
	Summary   GetTaskListSummaryItemList
}

type GetTaskListSummaryItem struct {
	Category    goaliyun.String
	SummaryName goaliyun.String
	Content     goaliyun.String
}

type GetTaskListContact struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
	JobId       goaliyun.String
}

type GetTaskListTaskList []GetTaskListTask

func (list *GetTaskListTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetTaskListTask)
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

type GetTaskListConversationDetailList []GetTaskListConversationDetail

func (list *GetTaskListConversationDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetTaskListConversationDetail)
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

type GetTaskListSummaryItemList []GetTaskListSummaryItem

func (list *GetTaskListSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetTaskListSummaryItem)
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
