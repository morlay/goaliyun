package hpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstancePasswordRequest struct {
	TOKEN       string `position:"Query" name:"TOKEN"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NewPassword string `position:"Query" name:"NewPassword"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstancePasswordRequest) Invoke(client goaliyun.Client) (*ModifyInstancePasswordResponse, error) {
	resp := &ModifyInstancePasswordResponse{}
	err := client.Request("hpc", "ModifyInstancePassword", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstancePasswordResponse struct {
	RequestId goaliyun.String
}
