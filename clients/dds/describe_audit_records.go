package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAuditRecordsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	QueryKeywords        string `position:"Query" name:"QueryKeywords"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Database             string `position:"Query" name:"Database"`
	Form                 string `position:"Query" name:"Form"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	NodeId               string `position:"Query" name:"NodeId"`
	User                 string `position:"Query" name:"User"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAuditRecordsRequest) Invoke(client goaliyun.Client) (*DescribeAuditRecordsResponse, error) {
	resp := &DescribeAuditRecordsResponse{}
	err := client.Request("dds", "DescribeAuditRecords", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAuditRecordsResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeAuditRecordsSQLRecordList
}

type DescribeAuditRecordsSQLRecord struct {
	DBName              goaliyun.String
	AccountName         goaliyun.String
	HostAddress         goaliyun.String
	Syntax              goaliyun.String
	TotalExecutionTimes goaliyun.Integer
	ReturnRowCounts     goaliyun.Integer
	ExecuteTime         goaliyun.String
	ThreadID            goaliyun.String
}

type DescribeAuditRecordsSQLRecordList []DescribeAuditRecordsSQLRecord

func (list *DescribeAuditRecordsSQLRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuditRecordsSQLRecord)
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
