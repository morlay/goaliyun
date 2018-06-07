package cms

import (
	"github.com/morlay/goaliyun"
)

type CreateAlarmRequest struct {
	Period             int64  `position:"Query" name:"Period"`
	Webhook            string `position:"Query" name:"Webhook"`
	ContactGroups      string `position:"Query" name:"ContactGroups"`
	EndTime            int64  `position:"Query" name:"EndTime"`
	Threshold          string `position:"Query" name:"Threshold"`
	StartTime          int64  `position:"Query" name:"StartTime"`
	Namespace          string `position:"Query" name:"Namespace"`
	Name               string `position:"Query" name:"Name"`
	EvaluationCount    int64  `position:"Query" name:"EvaluationCount"`
	SilenceTime        int64  `position:"Query" name:"SilenceTime"`
	MetricName         string `position:"Query" name:"MetricName"`
	NotifyType         int64  `position:"Query" name:"NotifyType"`
	ComparisonOperator string `position:"Query" name:"ComparisonOperator"`
	Dimensions         string `position:"Query" name:"Dimensions"`
	Statistics         string `position:"Query" name:"Statistics"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *CreateAlarmRequest) Invoke(client goaliyun.Client) (*CreateAlarmResponse, error) {
	resp := &CreateAlarmResponse{}
	err := client.Request("cms", "CreateAlarm", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAlarmResponse struct {
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
}
