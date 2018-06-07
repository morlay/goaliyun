package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyCommandRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	WorkingDir           string `position:"Query" name:"WorkingDir"`
	Description          string `position:"Query" name:"Description"`
	CommandId            string `position:"Query" name:"CommandId"`
	CommandContent       string `position:"Query" name:"CommandContent"`
	Timeout              int64  `position:"Query" name:"Timeout"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyCommandRequest) Invoke(client goaliyun.Client) (*ModifyCommandResponse, error) {
	resp := &ModifyCommandResponse{}
	err := client.Request("ecs", "ModifyCommand", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCommandResponse struct {
	RequestId goaliyun.String
}
