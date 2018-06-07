package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRecordLogsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeRecordLogsRequest) Invoke(client goaliyun.Client) (*DescribeRecordLogsResponse, error) {
	resp := &DescribeRecordLogsResponse{}
	err := client.Request("alidns", "DescribeRecordLogs", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRecordLogsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	RecordLogs DescribeRecordLogsRecordLogList
}

type DescribeRecordLogsRecordLog struct {
	ActionTime goaliyun.String
	Action     goaliyun.String
	Message    goaliyun.String
	ClientIp   goaliyun.String
}

type DescribeRecordLogsRecordLogList []DescribeRecordLogsRecordLog

func (list *DescribeRecordLogsRecordLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecordLogsRecordLog)
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
