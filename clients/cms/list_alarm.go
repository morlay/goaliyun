package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAlarmRequest struct {
	IsEnable   string `position:"Query" name:"IsEnable"`
	Name       string `position:"Query" name:"Name"`
	Namespace  string `position:"Query" name:"Namespace"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	Id         string `position:"Query" name:"Id"`
	State      string `position:"Query" name:"State"`
	Dimension  string `position:"Query" name:"Dimension"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListAlarmRequest) Invoke(client goaliyun.Client) (*ListAlarmResponse, error) {
	resp := &ListAlarmResponse{}
	err := client.Request("cms", "ListAlarm", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAlarmResponse struct {
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	NextToken goaliyun.Integer
	Total     goaliyun.Integer
	RequestId goaliyun.String
	AlarmList ListAlarmAlarmList
}

type ListAlarmAlarm struct {
	Id                 goaliyun.String
	Name               goaliyun.String
	Namespace          goaliyun.String
	MetricName         goaliyun.String
	Dimensions         goaliyun.String
	Period             goaliyun.Integer
	Statistics         goaliyun.String
	ComparisonOperator goaliyun.String
	Threshold          goaliyun.String
	EvaluationCount    goaliyun.Integer
	StartTime          goaliyun.Integer
	EndTime            goaliyun.Integer
	SilenceTime        goaliyun.Integer
	NotifyType         goaliyun.Integer
	Enable             bool
	State              goaliyun.String
	ContactGroups      goaliyun.String
	Webhook            goaliyun.String
}

type ListAlarmAlarmList []ListAlarmAlarm

func (list *ListAlarmAlarmList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlarmAlarm)
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
