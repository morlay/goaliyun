package vpc

import (
	"github.com/morlay/goaliyun"
)

type UnassociatePhysicalConnectionFromVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) Invoke(client goaliyun.Client) (*UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, error) {
	resp := &UnassociatePhysicalConnectionFromVirtualBorderRouterResponse{}
	err := client.Request("vpc", "UnassociatePhysicalConnectionFromVirtualBorderRouter", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnassociatePhysicalConnectionFromVirtualBorderRouterResponse struct {
	RequestId goaliyun.String
}
