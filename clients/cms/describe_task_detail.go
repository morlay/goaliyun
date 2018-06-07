package cms

import (
	"github.com/morlay/goaliyun"
)

type DescribeTaskDetailRequest struct {
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeTaskDetailRequest) Invoke(client goaliyun.Client) (*DescribeTaskDetailResponse, error) {
	resp := &DescribeTaskDetailResponse{}
	err := client.Request("cms", "DescribeTaskDetail", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTaskDetailResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
}
