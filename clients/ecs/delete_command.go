package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteCommandRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CommandId            string `position:"Query" name:"CommandId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteCommandRequest) Invoke(client goaliyun.Client) (*DeleteCommandResponse, error) {
	resp := &DeleteCommandResponse{}
	err := client.Request("ecs", "DeleteCommand", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCommandResponse struct {
	RequestId goaliyun.String
}
