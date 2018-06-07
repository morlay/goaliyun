package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSlowLogRecordsRequest struct {
	SQLId                int64  `position:"Query" name:"SQLId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	DBName               string `position:"Query" name:"DBName"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSlowLogRecordsRequest) Invoke(client goaliyun.Client) (*DescribeSlowLogRecordsResponse, error) {
	resp := &DescribeSlowLogRecordsResponse{}
	err := client.Request("rds", "DescribeSlowLogRecords", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSlowLogRecordsResponse struct {
	RequestId        goaliyun.String
	Engine           goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSlowLogRecordsSQLSlowRecordList
}

type DescribeSlowLogRecordsSQLSlowRecord struct {
	HostAddress        goaliyun.String
	DBName             goaliyun.String
	SQLText            goaliyun.String
	QueryTimes         goaliyun.Integer
	LockTimes          goaliyun.Integer
	ParseRowCounts     goaliyun.Integer
	ReturnRowCounts    goaliyun.Integer
	ExecutionStartTime goaliyun.String
}

type DescribeSlowLogRecordsSQLSlowRecordList []DescribeSlowLogRecordsSQLSlowRecord

func (list *DescribeSlowLogRecordsSQLSlowRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSlowLogRecordsSQLSlowRecord)
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
