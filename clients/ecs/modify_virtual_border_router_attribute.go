package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyVirtualBorderRouterAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CircuitCode          string `position:"Query" name:"CircuitCode"`
	VlanId               int64  `position:"Query" name:"VlanId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerGatewayIp        string `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask    string `position:"Query" name:"PeeringSubnetMask"`
	Name                 string `position:"Query" name:"Name"`
	LocalGatewayIp       string `position:"Query" name:"LocalGatewayIp"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVirtualBorderRouterAttributeRequest) Invoke(client goaliyun.Client) (*ModifyVirtualBorderRouterAttributeResponse, error) {
	resp := &ModifyVirtualBorderRouterAttributeResponse{}
	err := client.Request("ecs", "ModifyVirtualBorderRouterAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVirtualBorderRouterAttributeResponse struct {
	RequestId goaliyun.String
}
