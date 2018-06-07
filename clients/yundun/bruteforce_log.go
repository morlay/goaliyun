package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type BruteforceLogRequest struct {
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RecordType int64  `position:"Query" name:"RecordType"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *BruteforceLogRequest) Invoke(client goaliyun.Client) (*BruteforceLogResponse, error) {
	resp := &BruteforceLogResponse{}
	err := client.Request("yundun", "BruteforceLog", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BruteforceLogResponse struct {
	RequestId  goaliyun.String
	StartTime  goaliyun.String
	EndTime    goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	LogList    BruteforceLogBruteforceLogList
}

type BruteforceLogBruteforceLog struct {
	BlockTimes goaliyun.Integer
	SourceIp   goaliyun.String
	Status     goaliyun.Integer
	Time       goaliyun.String
	Location   goaliyun.String
}

type BruteforceLogBruteforceLogList []BruteforceLogBruteforceLog

func (list *BruteforceLogBruteforceLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BruteforceLogBruteforceLog)
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
