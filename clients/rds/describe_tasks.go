package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTasksRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	TaskAction           string `position:"Query" name:"TaskAction"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTasksRequest) Invoke(client goaliyun.Client) (*DescribeTasksResponse, error) {
	resp := &DescribeTasksResponse{}
	err := client.Request("rds", "DescribeTasks", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTasksResponse struct {
	RequestId        goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeTasksTaskProgressInfoList
}

type DescribeTasksTaskProgressInfo struct {
	DBName             goaliyun.String
	BeginTime          goaliyun.String
	ProgressInfo       goaliyun.String
	FinishTime         goaliyun.String
	TaskAction         goaliyun.String
	TaskId             goaliyun.String
	Progress           goaliyun.String
	ExpectedFinishTime goaliyun.String
	Status             goaliyun.String
	TaskErrorCode      goaliyun.String
	TaskErrorMessage   goaliyun.String
}

type DescribeTasksTaskProgressInfoList []DescribeTasksTaskProgressInfo

func (list *DescribeTasksTaskProgressInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTasksTaskProgressInfo)
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
