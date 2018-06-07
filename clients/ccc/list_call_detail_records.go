package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListCallDetailRecordsRequest struct {
	InstanceId         string `position:"Query" name:"InstanceId"`
	ContactDisposition string `position:"Query" name:"ContactDisposition"`
	ContactType        string `position:"Query" name:"ContactType"`
	Criteria           string `position:"Query" name:"Criteria"`
	PhoneNumber        string `position:"Query" name:"PhoneNumber"`
	PageSize           int64  `position:"Query" name:"PageSize"`
	OrderBy            string `position:"Query" name:"OrderBy"`
	StopTime           int64  `position:"Query" name:"StopTime"`
	StartTime          int64  `position:"Query" name:"StartTime"`
	PageNumber         int64  `position:"Query" name:"PageNumber"`
	WithRecording      string `position:"Query" name:"WithRecording"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *ListCallDetailRecordsRequest) Invoke(client goaliyun.Client) (*ListCallDetailRecordsResponse, error) {
	resp := &ListCallDetailRecordsResponse{}
	err := client.Request("ccc", "ListCallDetailRecords", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListCallDetailRecordsResponse struct {
	RequestId         goaliyun.String
	Success           bool
	Code              goaliyun.String
	Message           goaliyun.String
	HttpStatusCode    goaliyun.Integer
	CallDetailRecords ListCallDetailRecordsCallDetailRecords
}

type ListCallDetailRecordsCallDetailRecords struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       ListCallDetailRecordsCallDetailRecordList
}

type ListCallDetailRecordsCallDetailRecord struct {
	ContactId          goaliyun.String
	StartTime          goaliyun.Integer
	Duration           goaliyun.Integer
	Satisfaction       goaliyun.Integer
	ContactType        goaliyun.String
	ContactDisposition goaliyun.String
	CallingNumber      goaliyun.String
	CalledNumber       goaliyun.String
	AgentNames         goaliyun.String
	SkillGroupNames    goaliyun.String
	InstanceId         goaliyun.String
	ExtraAttr          goaliyun.String
	Agents             ListCallDetailRecordsCallDetailAgentList
	Recordings         ListCallDetailRecordsRecordingList
}

type ListCallDetailRecordsCallDetailAgent struct {
	ContactId      goaliyun.String
	AgentId        goaliyun.String
	AgentName      goaliyun.String
	SkillGroupName goaliyun.String
	QueueTime      goaliyun.Integer
	RingTime       goaliyun.Integer
	StartTime      goaliyun.Integer
	TalkTime       goaliyun.Integer
	HoldTime       goaliyun.Integer
	WorkTime       goaliyun.Integer
}

type ListCallDetailRecordsRecording struct {
	ContactId       goaliyun.String
	ContactType     goaliyun.String
	AgentId         goaliyun.String
	AgentName       goaliyun.String
	CallingNumber   goaliyun.String
	CalledNumber    goaliyun.String
	StartTime       goaliyun.Integer
	Duration        goaliyun.Integer
	FileName        goaliyun.String
	FilePath        goaliyun.String
	FileDescription goaliyun.String
	Channel         goaliyun.String
	InstanceId      goaliyun.String
}

type ListCallDetailRecordsCallDetailRecordList []ListCallDetailRecordsCallDetailRecord

func (list *ListCallDetailRecordsCallDetailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCallDetailRecordsCallDetailRecord)
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

type ListCallDetailRecordsCallDetailAgentList []ListCallDetailRecordsCallDetailAgent

func (list *ListCallDetailRecordsCallDetailAgentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCallDetailRecordsCallDetailAgent)
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

type ListCallDetailRecordsRecordingList []ListCallDetailRecordsRecording

func (list *ListCallDetailRecordsRecordingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCallDetailRecordsRecording)
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
