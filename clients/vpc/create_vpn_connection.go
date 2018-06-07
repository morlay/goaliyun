package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateVpnConnectionRequest struct {
	IkeConfig            string `position:"Query" name:"IkeConfig"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RemoteSubnet         string `position:"Query" name:"RemoteSubnet"`
	EffectImmediately    string `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpsecConfig          string `position:"Query" name:"IpsecConfig"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateVpnConnectionRequest) Invoke(client goaliyun.Client) (*CreateVpnConnectionResponse, error) {
	resp := &CreateVpnConnectionResponse{}
	err := client.Request("vpc", "CreateVpnConnection", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateVpnConnectionResponse struct {
	RequestId       goaliyun.String
	VpnConnectionId goaliyun.String
	Name            goaliyun.String
	CreateTime      goaliyun.Integer
}
