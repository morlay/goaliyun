package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListActiveAlertRuleRequest struct {
	Product  string `position:"Query" name:"Product"`
	UserId   string `position:"Query" name:"UserId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListActiveAlertRuleRequest) Invoke(client goaliyun.Client) (*ListActiveAlertRuleResponse, error) {
	resp := &ListActiveAlertRuleResponse{}
	err := client.Request("cms", "ListActiveAlertRule", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListActiveAlertRuleResponse struct {
	Success    bool
	Code       goaliyun.String
	Message    goaliyun.String
	RequestId  goaliyun.String
	Datapoints ListActiveAlertRuleAlarmList
}

type ListActiveAlertRuleAlarm struct {
	Uuid               goaliyun.String
	Name               goaliyun.String
	Namespace          goaliyun.String
	MetricName         goaliyun.String
	Period             goaliyun.String
	Statistics         goaliyun.String
	ComparisonOperator goaliyun.String
	Threshold          goaliyun.String
	EvaluationCount    goaliyun.String
	StartTime          goaliyun.String
	EndTime            goaliyun.String
	SilenceTime        goaliyun.String
	NotifyType         goaliyun.String
	Enable             goaliyun.String
	State              goaliyun.String
	ContactGroups      goaliyun.String
	Webhook            goaliyun.String
	RuleName           goaliyun.String
}

type ListActiveAlertRuleAlarmList []ListActiveAlertRuleAlarm

func (list *ListActiveAlertRuleAlarmList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListActiveAlertRuleAlarm)
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
