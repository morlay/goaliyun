package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobsByGroupRequest struct {
	InstanceId       string `position:"Query" name:"InstanceId"`
	JobFailureReason string `position:"Query" name:"JobFailureReason"`
	JobStatus        string `position:"Query" name:"JobStatus"`
	JobGroupId       string `position:"Query" name:"JobGroupId"`
	PageSize         int64  `position:"Query" name:"PageSize"`
	PageNumber       int64  `position:"Query" name:"PageNumber"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *ListJobsByGroupRequest) Invoke(client goaliyun.Client) (*ListJobsByGroupResponse, error) {
	resp := &ListJobsByGroupResponse{}
	err := client.Request("ccc", "ListJobsByGroup", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobsByGroupResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Jobs           ListJobsByGroupJobs
}

type ListJobsByGroupJobs struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       ListJobsByGroupJobList
}

type ListJobsByGroupJob struct {
	JobId          goaliyun.String
	GroupId        goaliyun.String
	ScenarioId     goaliyun.String
	StrategyId     goaliyun.String
	Priority       goaliyun.Integer
	SystemPriority goaliyun.Integer
	Status         goaliyun.String
	ReferenceId    goaliyun.String
	FailureReason  goaliyun.String
	Contacts       ListJobsByGroupContactList
	Extras         ListJobsByGroupKeyValuePairList
	Summary        ListJobsByGroupSummaryItemList
	CallingNumbers ListJobsByGroupCallingNumberList
}

type ListJobsByGroupContact struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
	JobId       goaliyun.String
}

type ListJobsByGroupKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type ListJobsByGroupSummaryItem struct {
	SummaryId            goaliyun.String
	GroupId              goaliyun.String
	JobId                goaliyun.String
	TaskId               goaliyun.String
	ConversationDetailId goaliyun.String
	Category             goaliyun.String
	SummaryName          goaliyun.String
	Content              goaliyun.String
}

type ListJobsByGroupJobList []ListJobsByGroupJob

func (list *ListJobsByGroupJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobsByGroupJob)
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

type ListJobsByGroupContactList []ListJobsByGroupContact

func (list *ListJobsByGroupContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobsByGroupContact)
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

type ListJobsByGroupKeyValuePairList []ListJobsByGroupKeyValuePair

func (list *ListJobsByGroupKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobsByGroupKeyValuePair)
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

type ListJobsByGroupSummaryItemList []ListJobsByGroupSummaryItem

func (list *ListJobsByGroupSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobsByGroupSummaryItem)
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

type ListJobsByGroupCallingNumberList []goaliyun.String

func (list *ListJobsByGroupCallingNumberList) UnmarshalJSON(data []byte) error {
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
