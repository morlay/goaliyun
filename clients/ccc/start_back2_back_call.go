package ccc

import (
	"github.com/morlay/goaliyun"
)

type StartBack2BackCallRequest struct {
	Caller           string `position:"Query" name:"Caller"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	CallCenterNumber string `position:"Query" name:"CallCenterNumber"`
	Callee           string `position:"Query" name:"Callee"`
	WorkflowId       string `position:"Query" name:"WorkflowId"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *StartBack2BackCallRequest) Invoke(client goaliyun.Client) (*StartBack2BackCallResponse, error) {
	resp := &StartBack2BackCallResponse{}
	err := client.Request("ccc", "StartBack2BackCall", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartBack2BackCallResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	StatusCode     goaliyun.String
	StatusDesc     goaliyun.String
	TaskId         goaliyun.String
	TimeStamp      goaliyun.String
}
