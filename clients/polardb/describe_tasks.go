package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTasksRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTasksRequest) Invoke(client goaliyun.Client) (*DescribeTasksResponse, error) {
	resp := &DescribeTasksResponse{}
	err := client.Request("polardb", "DescribeTasks", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTasksResponse struct {
	RequestId        goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Tasks            DescribeTasksTaskList
}

type DescribeTasksTask struct {
	TaskId             goaliyun.String
	BeginTime          goaliyun.String
	FinishTime         goaliyun.String
	ExpectedFinishTime goaliyun.String
	TaskAction         goaliyun.String
	Progress           goaliyun.Integer
	DBName             goaliyun.String
	ProgressInfo       goaliyun.String
	TaskErrorCode      goaliyun.String
	TaskErrorMessage   goaliyun.String
}

type DescribeTasksTaskList []DescribeTasksTask

func (list *DescribeTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTasksTask)
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
