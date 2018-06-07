package cms

import (
	"github.com/morlay/goaliyun"
)

type CreateTaskRequest struct {
	TaskType  string `position:"Query" name:"TaskType"`
	Address   string `position:"Query" name:"Address"`
	IspCity   string `position:"Query" name:"IspCity"`
	Options   string `position:"Query" name:"Options"`
	TaskName  string `position:"Query" name:"TaskName"`
	Interval  string `position:"Query" name:"Interval"`
	AlertRule string `position:"Query" name:"AlertRule"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *CreateTaskRequest) Invoke(client goaliyun.Client) (*CreateTaskResponse, error) {
	resp := &CreateTaskResponse{}
	err := client.Request("cms", "CreateTask", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateTaskResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
	AlertRule goaliyun.String
}
