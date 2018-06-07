package hpc

import (
	"github.com/morlay/goaliyun"
)

type RebootJumpserverRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int64  `position:"Query" name:"Force"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RebootJumpserverRequest) Invoke(client goaliyun.Client) (*RebootJumpserverResponse, error) {
	resp := &RebootJumpserverResponse{}
	err := client.Request("hpc", "RebootJumpserver", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RebootJumpserverResponse struct {
	RequestId goaliyun.String
}
