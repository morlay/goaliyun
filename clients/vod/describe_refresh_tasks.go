package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRefreshTasksRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ObjectPath           string `position:"Query" name:"ObjectPath"`
	DomainName           string `position:"Query" name:"DomainName"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ObjectType           string `position:"Query" name:"ObjectType"`
	TaskId               string `position:"Query" name:"TaskId"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRefreshTasksRequest) Invoke(client goaliyun.Client) (*DescribeRefreshTasksResponse, error) {
	resp := &DescribeRefreshTasksResponse{}
	err := client.Request("vod", "DescribeRefreshTasks", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRefreshTasksResponse struct {
	RequestId  goaliyun.String
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	TotalCount goaliyun.Integer
	Tasks      DescribeRefreshTasksTaskList
}

type DescribeRefreshTasksTask struct {
	TaskId       goaliyun.String
	ObjectPath   goaliyun.String
	Status       goaliyun.String
	Process      goaliyun.String
	ObjectType   goaliyun.String
	CreationTime goaliyun.String
	Description  goaliyun.String
}

type DescribeRefreshTasksTaskList []DescribeRefreshTasksTask

func (list *DescribeRefreshTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRefreshTasksTask)
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
