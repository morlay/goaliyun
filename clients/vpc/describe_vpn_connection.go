package vpc

import (
	"github.com/morlay/goaliyun"
)

type DescribeVpnConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpnConnectionRequest) Invoke(client goaliyun.Client) (*DescribeVpnConnectionResponse, error) {
	resp := &DescribeVpnConnectionResponse{}
	err := client.Request("vpc", "DescribeVpnConnection", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpnConnectionResponse struct {
	RequestId         goaliyun.String
	VpnConnectionId   goaliyun.String
	CustomerGatewayId goaliyun.String
	VpnGatewayId      goaliyun.String
	Name              goaliyun.String
	LocalSubnet       goaliyun.String
	RemoteSubnet      goaliyun.String
	CreateTime        goaliyun.Integer
	EffectImmediately bool
	Status            goaliyun.String
	IkeConfig         DescribeVpnConnectionIkeConfig
	IpsecConfig       DescribeVpnConnectionIpsecConfig
}

type DescribeVpnConnectionIkeConfig struct {
	Psk         goaliyun.String
	IkeVersion  goaliyun.String
	IkeMode     goaliyun.String
	IkeEncAlg   goaliyun.String
	IkeAuthAlg  goaliyun.String
	IkePfs      goaliyun.String
	IkeLifetime goaliyun.Integer
	LocalId     goaliyun.String
	RemoteId    goaliyun.String
}

type DescribeVpnConnectionIpsecConfig struct {
	IpsecEncAlg   goaliyun.String
	IpsecAuthAlg  goaliyun.String
	IpsecPfs      goaliyun.String
	IpsecLifetime goaliyun.Integer
}
