package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAlarmRulesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryAlarmRulesRequest) Invoke(client goaliyun.Client) (*QueryAlarmRulesResponse, error) {
	resp := &QueryAlarmRulesResponse{}
	err := client.Request("emr", "QueryAlarmRules", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAlarmRulesResponse struct {
	RequestId goaliyun.String
	AlarmList QueryAlarmRulesAlarmList
}

type QueryAlarmRulesAlarm struct {
	Name               goaliyun.String
	MetricName         goaliyun.String
	Period             goaliyun.Integer
	Threshold          goaliyun.String
	EvaluationCount    goaliyun.Integer
	StartTime          goaliyun.Integer
	EndTime            goaliyun.Integer
	SilenceTime        goaliyun.Integer
	NotifyType         goaliyun.Integer
	Enable             bool
	State              goaliyun.String
	ComparisonOperator goaliyun.String
	ContactGroups      goaliyun.String
}

type QueryAlarmRulesAlarmList []QueryAlarmRulesAlarm

func (list *QueryAlarmRulesAlarmList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAlarmRulesAlarm)
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
