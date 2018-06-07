package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPushListRequest struct {
	PageSize  int64  `position:"Query" name:"PageSize"`
	EndTime   string `position:"Query" name:"EndTime"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	StartTime string `position:"Query" name:"StartTime"`
	Page      int64  `position:"Query" name:"Page"`
	PushType  string `position:"Query" name:"PushType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QueryPushListRequest) Invoke(client goaliyun.Client) (*QueryPushListResponse, error) {
	resp := &QueryPushListResponse{}
	err := client.Request("push", "QueryPushList", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPushListResponse struct {
	RequestId        goaliyun.String
	HasNext          bool
	Page             goaliyun.Integer
	PageSize         goaliyun.Integer
	PushMessageInfos QueryPushListPushMessageInfoList
}

type QueryPushListPushMessageInfo struct {
	AppKey     goaliyun.Integer
	AppName    goaliyun.String
	MessageId  goaliyun.String
	Type       goaliyun.String
	DeviceType goaliyun.String
	PushTime   goaliyun.String
	Title      goaliyun.String
	Body       goaliyun.String
}

type QueryPushListPushMessageInfoList []QueryPushListPushMessageInfo

func (list *QueryPushListPushMessageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPushListPushMessageInfo)
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
