package cms

import (
	"github.com/morlay/goaliyun"
)

type StopTasksRequest struct {
	TaskIds  string `position:"Query" name:"TaskIds"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *StopTasksRequest) Invoke(client goaliyun.Client) (*StopTasksResponse, error) {
	resp := &StopTasksResponse{}
	err := client.Request("cms", "StopTasks", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopTasksResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
}
