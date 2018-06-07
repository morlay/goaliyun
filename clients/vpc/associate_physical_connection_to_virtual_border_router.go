package vpc

import (
	"github.com/morlay/goaliyun"
)

type AssociatePhysicalConnectionToVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CircuitCode          string `position:"Query" name:"CircuitCode"`
	VlanId               string `position:"Query" name:"VlanId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerGatewayIp        string `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask    string `position:"Query" name:"PeeringSubnetMask"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	LocalGatewayIp       string `position:"Query" name:"LocalGatewayIp"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AssociatePhysicalConnectionToVirtualBorderRouterRequest) Invoke(client goaliyun.Client) (*AssociatePhysicalConnectionToVirtualBorderRouterResponse, error) {
	resp := &AssociatePhysicalConnectionToVirtualBorderRouterResponse{}
	err := client.Request("vpc", "AssociatePhysicalConnectionToVirtualBorderRouter", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssociatePhysicalConnectionToVirtualBorderRouterResponse struct {
	RequestId goaliyun.String
}
