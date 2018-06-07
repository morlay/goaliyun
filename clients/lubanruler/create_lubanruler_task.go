package lubanruler

import (
	"github.com/morlay/goaliyun"
)

type CreateLubanrulerTaskRequest struct {
	TaskDO   string `position:"Body" name:"TaskDO"`
	Token    string `position:"Body" name:"Token"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateLubanrulerTaskRequest) Invoke(client goaliyun.Client) (*CreateLubanrulerTaskResponse, error) {
	resp := &CreateLubanrulerTaskResponse{}
	err := client.Request("lubanruler", "CreateLubanrulerTask", "2017-12-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateLubanrulerTaskResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
