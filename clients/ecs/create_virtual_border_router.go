package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CircuitCode          string `position:"Query" name:"CircuitCode"`
	VlanId               int64  `position:"Query" name:"VlanId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerGatewayIp        string `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask    string `position:"Query" name:"PeeringSubnetMask"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	Name                 string `position:"Query" name:"Name"`
	LocalGatewayIp       string `position:"Query" name:"LocalGatewayIp"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	VbrOwnerId           int64  `position:"Query" name:"VbrOwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateVirtualBorderRouterRequest) Invoke(client goaliyun.Client) (*CreateVirtualBorderRouterResponse, error) {
	resp := &CreateVirtualBorderRouterResponse{}
	err := client.Request("ecs", "CreateVirtualBorderRouter", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateVirtualBorderRouterResponse struct {
	RequestId goaliyun.String
	VbrId     goaliyun.String
}
