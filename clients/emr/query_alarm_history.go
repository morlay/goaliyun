package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAlarmHistoryRequest struct {
	Cursor          string `position:"Query" name:"Cursor"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Size            int64  `position:"Query" name:"Size"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	StartTimeStamp  int64  `position:"Query" name:"StartTimeStamp"`
	EndTimeStamp    int64  `position:"Query" name:"EndTimeStamp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryAlarmHistoryRequest) Invoke(client goaliyun.Client) (*QueryAlarmHistoryResponse, error) {
	resp := &QueryAlarmHistoryResponse{}
	err := client.Request("emr", "QueryAlarmHistory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAlarmHistoryResponse struct {
	RequestId        goaliyun.String
	Total            goaliyun.String
	Cursor           goaliyun.String
	AlarmHistoryList QueryAlarmHistoryEmrAlarmHistoryList
}

type QueryAlarmHistoryEmrAlarmHistory struct {
	ClusterId     goaliyun.String
	Role          goaliyun.String
	InstanceId    goaliyun.String
	Name          goaliyun.String
	MetricName    goaliyun.String
	AlarmTime     goaliyun.Integer
	LastTime      goaliyun.Integer
	State         goaliyun.String
	Status        goaliyun.Integer
	ContactGroups goaliyun.String
}

type QueryAlarmHistoryEmrAlarmHistoryList []QueryAlarmHistoryEmrAlarmHistory

func (list *QueryAlarmHistoryEmrAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAlarmHistoryEmrAlarmHistory)
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
