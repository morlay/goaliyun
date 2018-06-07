package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyVpnConnectionAttributeRequest struct {
	IkeConfig            string `position:"Query" name:"IkeConfig"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RemoteSubnet         string `position:"Query" name:"RemoteSubnet"`
	EffectImmediately    string `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpsecConfig          string `position:"Query" name:"IpsecConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVpnConnectionAttributeRequest) Invoke(client goaliyun.Client) (*ModifyVpnConnectionAttributeResponse, error) {
	resp := &ModifyVpnConnectionAttributeResponse{}
	err := client.Request("vpc", "ModifyVpnConnectionAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVpnConnectionAttributeResponse struct {
	RequestId         goaliyun.String
	VpnConnectionId   goaliyun.String
	CustomerGatewayId goaliyun.String
	VpnGatewayId      goaliyun.String
	Name              goaliyun.String
	Description       goaliyun.String
	LocalSubnet       goaliyun.String
	RemoteSubnet      goaliyun.String
	CreateTime        goaliyun.Integer
	EffectImmediately bool
	IkeConfig         ModifyVpnConnectionAttributeIkeConfig
	IpsecConfig       ModifyVpnConnectionAttributeIpsecConfig
}

type ModifyVpnConnectionAttributeIkeConfig struct {
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

type ModifyVpnConnectionAttributeIpsecConfig struct {
	IpsecEncAlg   goaliyun.String
	IpsecAuthAlg  goaliyun.String
	IpsecPfs      goaliyun.String
	IpsecLifetime goaliyun.Integer
}
