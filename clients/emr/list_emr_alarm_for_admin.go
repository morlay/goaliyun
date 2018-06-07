package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListEmrAlarmForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	UniqueKey       string `position:"Query" name:"UniqueKey"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	Topic           string `position:"Query" name:"Topic"`
	EventType       string `position:"Query" name:"EventType"`
	EntityId        string `position:"Query" name:"EntityId"`
	Priority        int64  `position:"Query" name:"Priority"`
	ClusterBizId    string `position:"Query" name:"ClusterBizId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListEmrAlarmForAdminRequest) Invoke(client goaliyun.Client) (*ListEmrAlarmForAdminResponse, error) {
	resp := &ListEmrAlarmForAdminResponse{}
	err := client.Request("emr", "ListEmrAlarmForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListEmrAlarmForAdminResponse struct {
	RequestId        goaliyun.String
	TotalCount       goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageSize         goaliyun.Integer
	AlarmHistoryList ListEmrAlarmForAdminAlarmHistoryList
}

type ListEmrAlarmForAdminAlarmHistory struct {
	Id           goaliyun.Integer
	EventType    goaliyun.String
	Topic        goaliyun.String
	UniqueKey    goaliyun.String
	EntityId     goaliyun.String
	Priority     goaliyun.Integer
	Body         goaliyun.String
	Status       goaliyun.String
	ClusterBizId goaliyun.String
	GmtCreate    goaliyun.String
	GmtModified  goaliyun.String
}

type ListEmrAlarmForAdminAlarmHistoryList []ListEmrAlarmForAdminAlarmHistory

func (list *ListEmrAlarmForAdminAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEmrAlarmForAdminAlarmHistory)
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
