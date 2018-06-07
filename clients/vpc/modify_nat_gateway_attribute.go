package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyNatGatewayAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyNatGatewayAttributeRequest) Invoke(client goaliyun.Client) (*ModifyNatGatewayAttributeResponse, error) {
	resp := &ModifyNatGatewayAttributeResponse{}
	err := client.Request("vpc", "ModifyNatGatewayAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyNatGatewayAttributeResponse struct {
	RequestId goaliyun.String
}
