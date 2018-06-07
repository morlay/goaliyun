package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPushStatByAppRequest struct {
	Granularity string `position:"Query" name:"Granularity"`
	EndTime     string `position:"Query" name:"EndTime"`
	AppKey      int64  `position:"Query" name:"AppKey"`
	StartTime   string `position:"Query" name:"StartTime"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *QueryPushStatByAppRequest) Invoke(client goaliyun.Client) (*QueryPushStatByAppResponse, error) {
	resp := &QueryPushStatByAppResponse{}
	err := client.Request("push", "QueryPushStatByApp", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPushStatByAppResponse struct {
	RequestId    goaliyun.String
	AppPushStats QueryPushStatByAppAppPushStatList
}

type QueryPushStatByAppAppPushStat struct {
	Time                   goaliyun.String
	AcceptCount            goaliyun.Integer
	SentCount              goaliyun.Integer
	ReceivedCount          goaliyun.Integer
	OpenedCount            goaliyun.Integer
	DeletedCount           goaliyun.Integer
	SmsSentCount           goaliyun.Integer
	SmsSkipCount           goaliyun.Integer
	SmsFailedCount         goaliyun.Integer
	SmsReceiveSuccessCount goaliyun.Integer
	SmsReceiveFailedCount  goaliyun.Integer
}

type QueryPushStatByAppAppPushStatList []QueryPushStatByAppAppPushStat

func (list *QueryPushStatByAppAppPushStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushStatByAppAppPushStat)
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
