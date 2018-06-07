package cms

import (
	"github.com/morlay/goaliyun"
)

type DescribeTasksRequest struct {
	TaskType string `position:"Query" name:"TaskType"`
	PageSize int64  `position:"Query" name:"PageSize"`
	Page     int64  `position:"Query" name:"Page"`
	Keyword  string `position:"Query" name:"Keyword"`
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeTasksRequest) Invoke(client goaliyun.Client) (*DescribeTasksResponse, error) {
	resp := &DescribeTasksResponse{}
	err := client.Request("cms", "DescribeTasks", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTasksResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	Success    goaliyun.String
	RequestId  goaliyun.String
	Data       goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
}
