package pvtz

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeChangeLogsRequest struct {
	EntityType     string `position:"Query" name:"EntityType"`
	PageSize       int64  `position:"Query" name:"PageSize"`
	UserClientIp   string `position:"Query" name:"UserClientIp"`
	ZoneId         string `position:"Query" name:"ZoneId"`
	Keyword        string `position:"Query" name:"Keyword"`
	Lang           string `position:"Query" name:"Lang"`
	StartTimestamp int64  `position:"Query" name:"StartTimestamp"`
	PageNumber     int64  `position:"Query" name:"PageNumber"`
	EndTimestamp   int64  `position:"Query" name:"EndTimestamp"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeChangeLogsRequest) Invoke(client goaliyun.Client) (*DescribeChangeLogsResponse, error) {
	resp := &DescribeChangeLogsResponse{}
	err := client.Request("pvtz", "DescribeChangeLogs", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeChangeLogsResponse struct {
	RequestId  goaliyun.String
	TotalItems goaliyun.Integer
	TotalPages goaliyun.Integer
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	ChangeLogs DescribeChangeLogsChangeLogList
}

type DescribeChangeLogsChangeLog struct {
	OperTime      goaliyun.String
	OperAction    goaliyun.String
	OperObject    goaliyun.String
	EntityId      goaliyun.String
	EntityName    goaliyun.String
	OperIp        goaliyun.String
	OperTimestamp goaliyun.Integer
	Id            goaliyun.Integer
	Content       goaliyun.String
}

type DescribeChangeLogsChangeLogList []DescribeChangeLogsChangeLog

func (list *DescribeChangeLogsChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeChangeLogsChangeLog)
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
