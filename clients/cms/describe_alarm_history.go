package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAlarmHistoryRequest struct {
	AlertName  string `position:"Query" name:"AlertName"`
	GroupId    string `position:"Query" name:"GroupId"`
	EndTime    string `position:"Query" name:"EndTime"`
	RuleName   string `position:"Query" name:"RuleName"`
	StartTime  string `position:"Query" name:"StartTime"`
	Ascending  string `position:"Query" name:"Ascending"`
	OnlyCount  string `position:"Query" name:"OnlyCount"`
	Namespace  string `position:"Query" name:"Namespace"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	State      string `position:"Query" name:"State"`
	Page       int64  `position:"Query" name:"Page"`
	MetricName string `position:"Query" name:"MetricName"`
	Status     string `position:"Query" name:"Status"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeAlarmHistoryRequest) Invoke(client goaliyun.Client) (*DescribeAlarmHistoryResponse, error) {
	resp := &DescribeAlarmHistoryResponse{}
	err := client.Request("cms", "DescribeAlarmHistory", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAlarmHistoryResponse struct {
	Success          bool
	Code             goaliyun.String
	Message          goaliyun.String
	Total            goaliyun.String
	RequestId        goaliyun.String
	AlarmHistoryList DescribeAlarmHistoryAlarmHistoryList
}

type DescribeAlarmHistoryAlarmHistory struct {
	Id              goaliyun.String
	AlertName       goaliyun.String
	GroupId         goaliyun.String
	Namespace       goaliyun.String
	MetricName      goaliyun.String
	Dimensions      goaliyun.String
	Expression      goaliyun.String
	EvaluationCount goaliyun.Integer
	Value           goaliyun.String
	AlertTime       goaliyun.Integer
	LastTime        goaliyun.Integer
	Level           goaliyun.String
	PreLevel        goaliyun.String
	RuleName        goaliyun.String
	State           goaliyun.String
	Status          goaliyun.Integer
	UserId          goaliyun.String
	Webhooks        goaliyun.String
	InstanceName    goaliyun.String
	ContactGroups   DescribeAlarmHistoryContactGroupList
	Contacts        DescribeAlarmHistoryContactList
	ContactALIIMs   DescribeAlarmHistoryContactALIIMList
	ContactSmses    DescribeAlarmHistoryContactSmseList
	ContactMails    DescribeAlarmHistoryContactMailList
}

type DescribeAlarmHistoryAlarmHistoryList []DescribeAlarmHistoryAlarmHistory

func (list *DescribeAlarmHistoryAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAlarmHistoryAlarmHistory)
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

type DescribeAlarmHistoryContactGroupList []goaliyun.String

func (list *DescribeAlarmHistoryContactGroupList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactList []goaliyun.String

func (list *DescribeAlarmHistoryContactList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactALIIMList []goaliyun.String

func (list *DescribeAlarmHistoryContactALIIMList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactSmseList []goaliyun.String

func (list *DescribeAlarmHistoryContactSmseList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactMailList []goaliyun.String

func (list *DescribeAlarmHistoryContactMailList) UnmarshalJSON(data []byte) error {
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
