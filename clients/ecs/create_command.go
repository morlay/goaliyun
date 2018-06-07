package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateCommandRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	WorkingDir           string `position:"Query" name:"WorkingDir"`
	Description          string `position:"Query" name:"Description"`
	Type                 string `position:"Query" name:"Type"`
	CommandContent       string `position:"Query" name:"CommandContent"`
	Timeout              int64  `position:"Query" name:"Timeout"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateCommandRequest) Invoke(client goaliyun.Client) (*CreateCommandResponse, error) {
	resp := &CreateCommandResponse{}
	err := client.Request("ecs", "CreateCommand", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateCommandResponse struct {
	RequestId goaliyun.String
	CommandId goaliyun.String
}
