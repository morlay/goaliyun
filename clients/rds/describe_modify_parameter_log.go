package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeModifyParameterLogRequest struct {
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

func (req *DescribeModifyParameterLogRequest) Invoke(client goaliyun.Client) (*DescribeModifyParameterLogResponse, error) {
	resp := &DescribeModifyParameterLogResponse{}
	err := client.Request("rds", "DescribeModifyParameterLog", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeModifyParameterLogResponse struct {
	RequestId        goaliyun.String
	Engine           goaliyun.String
	DBInstanceId     goaliyun.String
	EngineVersion    goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeModifyParameterLogParameterChangeLogList
}

type DescribeModifyParameterLogParameterChangeLog struct {
	ModifyTime        goaliyun.String
	OldParameterValue goaliyun.String
	NewParameterValue goaliyun.String
	ParameterName     goaliyun.String
	Status            goaliyun.String
}

type DescribeModifyParameterLogParameterChangeLogList []DescribeModifyParameterLogParameterChangeLog

func (list *DescribeModifyParameterLogParameterChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeModifyParameterLogParameterChangeLog)
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
