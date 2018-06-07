package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRealtimeDiagnosesRequest struct {
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

func (req *DescribeRealtimeDiagnosesRequest) Invoke(client goaliyun.Client) (*DescribeRealtimeDiagnosesResponse, error) {
	resp := &DescribeRealtimeDiagnosesResponse{}
	err := client.Request("rds", "DescribeRealtimeDiagnoses", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRealtimeDiagnosesResponse struct {
	RequestId        goaliyun.String
	Engine           goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Tasks            DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList
}

type DescribeRealtimeDiagnosesRealtimeDiagnoseTasks struct {
	CreateTime  goaliyun.String
	TaskId      goaliyun.String
	HealthScore goaliyun.String
	Status      goaliyun.String
}

type DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList []DescribeRealtimeDiagnosesRealtimeDiagnoseTasks

func (list *DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRealtimeDiagnosesRealtimeDiagnoseTasks)
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
