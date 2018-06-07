package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifySslVpnServerRequest struct {
	Cipher               string `position:"Query" name:"Cipher"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientIpPool         string `position:"Query" name:"ClientIpPool"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Compress             string `position:"Query" name:"Compress"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	Port                 int64  `position:"Query" name:"Port"`
	Proto                string `position:"Query" name:"Proto"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySslVpnServerRequest) Invoke(client goaliyun.Client) (*ModifySslVpnServerResponse, error) {
	resp := &ModifySslVpnServerResponse{}
	err := client.Request("vpc", "ModifySslVpnServer", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySslVpnServerResponse struct {
	RequestId      goaliyun.String
	RegionId       goaliyun.String
	SslVpnServerId goaliyun.String
	VpnGatewayId   goaliyun.String
	Name           goaliyun.String
	LocalSubnet    goaliyun.String
	ClientIpPool   goaliyun.String
	CreateTime     goaliyun.Integer
	Cipher         goaliyun.String
	Proto          goaliyun.String
	Port           goaliyun.Integer
	Compress       bool
	Connections    goaliyun.Integer
	MaxConnections goaliyun.Integer
	InternetIp     goaliyun.String
}
