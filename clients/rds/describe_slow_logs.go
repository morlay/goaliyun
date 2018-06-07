package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSlowLogsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	SortKey              string `position:"Query" name:"SortKey"`
	DBName               string `position:"Query" name:"DBName"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSlowLogsRequest) Invoke(client goaliyun.Client) (*DescribeSlowLogsResponse, error) {
	resp := &DescribeSlowLogsResponse{}
	err := client.Request("rds", "DescribeSlowLogs", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSlowLogsResponse struct {
	RequestId        goaliyun.String
	Engine           goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSlowLogsSQLSlowLogList
}

type DescribeSlowLogsSQLSlowLog struct {
	SlowLogId                     goaliyun.Integer
	SQLId                         goaliyun.Integer
	DBName                        goaliyun.String
	SQLText                       goaliyun.String
	MySQLTotalExecutionCounts     goaliyun.Integer
	MySQLTotalExecutionTimes      goaliyun.Integer
	TotalLockTimes                goaliyun.Integer
	MaxLockTime                   goaliyun.Integer
	ParseTotalRowCounts           goaliyun.Integer
	ParseMaxRowCount              goaliyun.Integer
	ReturnTotalRowCounts          goaliyun.Integer
	ReturnMaxRowCount             goaliyun.Integer
	CreateTime                    goaliyun.String
	SQLServerTotalExecutionCounts goaliyun.Integer
	SQLServerTotalExecutionTimes  goaliyun.Integer
	TotalLogicalReadCounts        goaliyun.Integer
	TotalPhysicalReadCounts       goaliyun.Integer
	ReportTime                    goaliyun.String
	MaxExecutionTime              goaliyun.Integer
	AvgExecutionTime              goaliyun.Integer
}

type DescribeSlowLogsSQLSlowLogList []DescribeSlowLogsSQLSlowLog

func (list *DescribeSlowLogsSQLSlowLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSlowLogsSQLSlowLog)
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
