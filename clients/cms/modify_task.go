package cms

import (
	"github.com/morlay/goaliyun"
)

type ModifyTaskRequest struct {
	Address   string `position:"Query" name:"Address"`
	IspCity   string `position:"Query" name:"IspCity"`
	Options   string `position:"Query" name:"Options"`
	TaskName  string `position:"Query" name:"TaskName"`
	Interval  string `position:"Query" name:"Interval"`
	AlertRule string `position:"Query" name:"AlertRule"`
	TaskId    string `position:"Query" name:"TaskId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ModifyTaskRequest) Invoke(client goaliyun.Client) (*ModifyTaskResponse, error) {
	resp := &ModifyTaskResponse{}
	err := client.Request("cms", "ModifyTask", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyTaskResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   goaliyun.String
	RequestId goaliyun.String
	Data      goaliyun.String
}
