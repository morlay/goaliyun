package lubanruler

import (
	"github.com/morlay/goaliyun"
)

type UpdateLubanrulerTaskRequest struct {
	AoneInfo string `position:"Body" name:"AoneInfo"`
	Token    string `position:"Body" name:"Token"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateLubanrulerTaskRequest) Invoke(client goaliyun.Client) (*UpdateLubanrulerTaskResponse, error) {
	resp := &UpdateLubanrulerTaskResponse{}
	err := client.Request("lubanruler", "UpdateLubanrulerTask", "2017-12-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateLubanrulerTaskResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
