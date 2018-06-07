package hpc

import (
	"github.com/morlay/goaliyun"
)

type StopJumpserverRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int64  `position:"Query" name:"Force"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *StopJumpserverRequest) Invoke(client goaliyun.Client) (*StopJumpserverResponse, error) {
	resp := &StopJumpserverResponse{}
	err := client.Request("hpc", "StopJumpserver", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopJumpserverResponse struct {
	RequestId goaliyun.String
}
