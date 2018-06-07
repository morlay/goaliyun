package cms

import (
	"github.com/morlay/goaliyun"
)

type StartTasksRequest struct {
	TaskIds  string `position:"Query" name:"TaskIds"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *StartTasksRequest) Invoke(client goaliyun.Client) (*StartTasksResponse, error) {
	resp := &StartTasksResponse{}
	err := client.Request("cms", "StartTasks", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartTasksResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
}
