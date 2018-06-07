package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteVpnGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteVpnGatewayRequest) Invoke(client goaliyun.Client) (*DeleteVpnGatewayResponse, error) {
	resp := &DeleteVpnGatewayResponse{}
	err := client.Request("vpc", "DeleteVpnGateway", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVpnGatewayResponse struct {
	RequestId goaliyun.String
}
