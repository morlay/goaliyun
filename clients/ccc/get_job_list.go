package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetJobListRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	Status     int64  `position:"Query" name:"Status"`
	QueryAll   string `position:"Query" name:"QueryAll"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetJobListRequest) Invoke(client goaliyun.Client) (*GetJobListResponse, error) {
	resp := &GetJobListResponse{}
	err := client.Request("ccc", "GetJobList", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetJobListResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Jobs           GetJobListJobs
}

type GetJobListJobs struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       GetJobListJobList
}

type GetJobListJob struct {
	JobId          goaliyun.String
	GroupId        goaliyun.String
	ScenarioId     goaliyun.String
	StrategyId     goaliyun.String
	Priority       goaliyun.Integer
	SystemPriority goaliyun.Integer
	Status         goaliyun.String
	ReferenceId    goaliyun.String
	FailureReason  goaliyun.String
	Contacts       GetJobListContactList
	Extras         GetJobListKeyValuePairList
	Summary        GetJobListSummaryItemList
	CallingNumbers GetJobListCallingNumberList
}

type GetJobListContact struct {
	ContactId   goaliyun.String
	ContactName goaliyun.String
	Honorific   goaliyun.String
	Role        goaliyun.String
	PhoneNumber goaliyun.String
	State       goaliyun.String
	ReferenceId goaliyun.String
	JobId       goaliyun.String
}

type GetJobListKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type GetJobListSummaryItem struct {
	SummaryId            goaliyun.String
	GroupId              goaliyun.String
	JobId                goaliyun.String
	TaskId               goaliyun.String
	ConversationDetailId goaliyun.String
	Category             goaliyun.String
	SummaryName          goaliyun.String
	Content              goaliyun.String
}

type GetJobListJobList []GetJobListJob

func (list *GetJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobListJob)
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

type GetJobListContactList []GetJobListContact

func (list *GetJobListContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobListContact)
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

type GetJobListKeyValuePairList []GetJobListKeyValuePair

func (list *GetJobListKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobListKeyValuePair)
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

type GetJobListSummaryItemList []GetJobListSummaryItem

func (list *GetJobListSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobListSummaryItem)
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

type GetJobListCallingNumberList []goaliyun.String

func (list *GetJobListCallingNumberList) UnmarshalJSON(data []byte) error {
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
