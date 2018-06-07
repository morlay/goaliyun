package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyNatGatewaySpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Spec                 string `position:"Query" name:"Spec"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyNatGatewaySpecRequest) Invoke(client goaliyun.Client) (*ModifyNatGatewaySpecResponse, error) {
	resp := &ModifyNatGatewaySpecResponse{}
	err := client.Request("vpc", "ModifyNatGatewaySpec", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyNatGatewaySpecResponse struct {
	RequestId goaliyun.String
}
