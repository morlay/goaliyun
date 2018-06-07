package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type LogineventLogRequest struct {
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RecordType int64  `position:"Query" name:"RecordType"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *LogineventLogRequest) Invoke(client goaliyun.Client) (*LogineventLogResponse, error) {
	resp := &LogineventLogResponse{}
	err := client.Request("yundun", "LogineventLog", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type LogineventLogResponse struct {
	RequestId  goaliyun.String
	StartTime  goaliyun.String
	EndTime    goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	LogList    LogineventLogLoginEventLogList
}

type LogineventLogLoginEventLog struct {
	BlockTimes goaliyun.Integer
	SourceIp   goaliyun.String
	Status     goaliyun.Integer
	Time       goaliyun.String
	Location   goaliyun.String
}

type LogineventLogLoginEventLogList []LogineventLogLoginEventLog

func (list *LogineventLogLoginEventLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]LogineventLogLoginEventLog)
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
