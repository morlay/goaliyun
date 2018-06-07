package cms

import (
	"github.com/morlay/goaliyun"
)

type UpdateAlarmRequest struct {
	Period             int64  `position:"Query" name:"Period"`
	Webhook            string `position:"Query" name:"Webhook"`
	ContactGroups      string `position:"Query" name:"ContactGroups"`
	EndTime            int64  `position:"Query" name:"EndTime"`
	Threshold          string `position:"Query" name:"Threshold"`
	StartTime          int64  `position:"Query" name:"StartTime"`
	Name               string `position:"Query" name:"Name"`
	EvaluationCount    int64  `position:"Query" name:"EvaluationCount"`
	SilenceTime        int64  `position:"Query" name:"SilenceTime"`
	Id                 string `position:"Query" name:"Id"`
	NotifyType         int64  `position:"Query" name:"NotifyType"`
	ComparisonOperator string `position:"Query" name:"ComparisonOperator"`
	Statistics         string `position:"Query" name:"Statistics"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *UpdateAlarmRequest) Invoke(client goaliyun.Client) (*UpdateAlarmResponse, error) {
	resp := &UpdateAlarmResponse{}
	err := client.Request("cms", "UpdateAlarm", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateAlarmResponse struct {
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
