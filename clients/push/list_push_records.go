package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPushRecordsRequest struct {
	PageSize  int64  `position:"Query" name:"PageSize"`
	EndTime   string `position:"Query" name:"EndTime"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	StartTime string `position:"Query" name:"StartTime"`
	Page      int64  `position:"Query" name:"Page"`
	PushType  string `position:"Query" name:"PushType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListPushRecordsRequest) Invoke(client goaliyun.Client) (*ListPushRecordsResponse, error) {
	resp := &ListPushRecordsResponse{}
	err := client.Request("push", "ListPushRecords", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPushRecordsResponse struct {
	RequestId        goaliyun.String
	Total            goaliyun.Integer
	Page             goaliyun.Integer
	PageSize         goaliyun.Integer
	PushMessageInfos ListPushRecordsPushMessageInfoList
}

type ListPushRecordsPushMessageInfo struct {
	AppKey     goaliyun.Integer
	AppName    goaliyun.String
	MessageId  goaliyun.String
	Type       goaliyun.String
	DeviceType goaliyun.String
	PushTime   goaliyun.String
	Title      goaliyun.String
	Body       goaliyun.String
}

type ListPushRecordsPushMessageInfoList []ListPushRecordsPushMessageInfo

func (list *ListPushRecordsPushMessageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPushRecordsPushMessageInfo)
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
