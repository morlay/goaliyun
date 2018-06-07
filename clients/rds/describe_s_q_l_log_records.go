package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLLogRecordsRequest struct {
	SQLId                int64  `position:"Query" name:"SQLId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	QueryKeywords        string `position:"Query" name:"QueryKeywords"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Database             string `position:"Query" name:"Database"`
	Form                 string `position:"Query" name:"Form"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	User                 string `position:"Query" name:"User"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSQLLogRecordsRequest) Invoke(client goaliyun.Client) (*DescribeSQLLogRecordsResponse, error) {
	resp := &DescribeSQLLogRecordsResponse{}
	err := client.Request("rds", "DescribeSQLLogRecords", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLLogRecordsResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSQLLogRecordsSQLRecordList
}

type DescribeSQLLogRecordsSQLRecord struct {
	DBName              goaliyun.String
	AccountName         goaliyun.String
	HostAddress         goaliyun.String
	SQLText             goaliyun.String
	TotalExecutionTimes goaliyun.Integer
	ReturnRowCounts     goaliyun.Integer
	ExecuteTime         goaliyun.String
	ThreadID            goaliyun.String
}

type DescribeSQLLogRecordsSQLRecordList []DescribeSQLLogRecordsSQLRecord

func (list *DescribeSQLLogRecordsSQLRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogRecordsSQLRecord)
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
