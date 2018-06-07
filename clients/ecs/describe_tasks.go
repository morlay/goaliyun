package ecs

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
	TaskIds              string `position:"Query" name:"TaskIds"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	TaskStatus           string `position:"Query" name:"TaskStatus"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	TaskAction           string `position:"Query" name:"TaskAction"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTasksRequest) Invoke(client goaliyun.Client) (*DescribeTasksResponse, error) {
	resp := &DescribeTasksResponse{}
	err := client.Request("ecs", "DescribeTasks", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTasksResponse struct {
	RequestId  goaliyun.String
	RegionId   goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TaskSet    DescribeTasksTaskList
}

type DescribeTasksTask struct {
	TaskId        goaliyun.String
	TaskAction    goaliyun.String
	TaskStatus    goaliyun.String
	SupportCancel goaliyun.String
	CreationTime  goaliyun.String
	FinishedTime  goaliyun.String
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
