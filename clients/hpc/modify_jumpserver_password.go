package hpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyJumpserverPasswordRequest struct {
	TOKEN       string `position:"Query" name:"TOKEN"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NewPassword string `position:"Query" name:"NewPassword"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ModifyJumpserverPasswordRequest) Invoke(client goaliyun.Client) (*ModifyJumpserverPasswordResponse, error) {
	resp := &ModifyJumpserverPasswordResponse{}
	err := client.Request("hpc", "ModifyJumpserverPassword", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyJumpserverPasswordResponse struct {
	RequestId goaliyun.String
}
