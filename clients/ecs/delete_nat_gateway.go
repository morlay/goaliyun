package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteNatGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteNatGatewayRequest) Invoke(client goaliyun.Client) (*DeleteNatGatewayResponse, error) {
	resp := &DeleteNatGatewayResponse{}
	err := client.Request("ecs", "DeleteNatGateway", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNatGatewayResponse struct {
	RequestId goaliyun.String
}
