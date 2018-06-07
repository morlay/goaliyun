package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeErrorLogsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeErrorLogsRequest) Invoke(client goaliyun.Client) (*DescribeErrorLogsResponse, error) {
	resp := &DescribeErrorLogsResponse{}
	err := client.Request("rds", "DescribeErrorLogs", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeErrorLogsResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeErrorLogsErrorLogList
}

type DescribeErrorLogsErrorLog struct {
	ErrorInfo  goaliyun.String
	CreateTime goaliyun.String
}

type DescribeErrorLogsErrorLogList []DescribeErrorLogsErrorLog

func (list *DescribeErrorLogsErrorLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeErrorLogsErrorLog)
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
