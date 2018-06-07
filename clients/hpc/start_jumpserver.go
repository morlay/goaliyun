package hpc

import (
	"github.com/morlay/goaliyun"
)

type StartJumpserverRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int64  `position:"Query" name:"Force"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *StartJumpserverRequest) Invoke(client goaliyun.Client) (*StartJumpserverResponse, error) {
	resp := &StartJumpserverResponse{}
	err := client.Request("hpc", "StartJumpserver", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartJumpserverResponse struct {
	RequestId goaliyun.String
}
