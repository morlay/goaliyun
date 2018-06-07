package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPushStatByMsgRequest struct {
	MessageId int64  `position:"Query" name:"MessageId"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QueryPushStatByMsgRequest) Invoke(client goaliyun.Client) (*QueryPushStatByMsgResponse, error) {
	resp := &QueryPushStatByMsgResponse{}
	err := client.Request("push", "QueryPushStatByMsg", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPushStatByMsgResponse struct {
	RequestId goaliyun.String
	PushStats QueryPushStatByMsgPushStatList
}

type QueryPushStatByMsgPushStat struct {
	MessageId              goaliyun.String
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

type QueryPushStatByMsgPushStatList []QueryPushStatByMsgPushStat

func (list *QueryPushStatByMsgPushStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushStatByMsgPushStat)
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
