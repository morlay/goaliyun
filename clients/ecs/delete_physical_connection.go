package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeletePhysicalConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeletePhysicalConnectionRequest) Invoke(client goaliyun.Client) (*DeletePhysicalConnectionResponse, error) {
	resp := &DeletePhysicalConnectionResponse{}
	err := client.Request("ecs", "DeletePhysicalConnection", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePhysicalConnectionResponse struct {
	RequestId goaliyun.String
}
