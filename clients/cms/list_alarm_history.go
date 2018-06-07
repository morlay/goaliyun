package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAlarmHistoryRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	EndTime   string `position:"Query" name:"EndTime"`
	Id        string `position:"Query" name:"Id"`
	StartTime string `position:"Query" name:"StartTime"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListAlarmHistoryRequest) Invoke(client goaliyun.Client) (*ListAlarmHistoryResponse, error) {
	resp := &ListAlarmHistoryResponse{}
	err := client.Request("cms", "ListAlarmHistory", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAlarmHistoryResponse struct {
	Success          bool
	Code             goaliyun.String
	Message          goaliyun.String
	Cursor           goaliyun.String
	RequestId        goaliyun.String
	AlarmHistoryList ListAlarmHistoryAlarmHistoryList
}

type ListAlarmHistoryAlarmHistory struct {
	Id              goaliyun.String
	Name            goaliyun.String
	Namespace       goaliyun.String
	MetricName      goaliyun.String
	Dimension       goaliyun.String
	EvaluationCount goaliyun.Integer
	Value           goaliyun.String
	AlarmTime       goaliyun.Integer
	LastTime        goaliyun.Integer
	State           goaliyun.String
	Status          goaliyun.Integer
	ContactGroups   goaliyun.String
	InstanceName    goaliyun.String
}

type ListAlarmHistoryAlarmHistoryList []ListAlarmHistoryAlarmHistory

func (list *ListAlarmHistoryAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlarmHistoryAlarmHistory)
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
