package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateSslVpnServerRequest struct {
	Cipher               string `position:"Query" name:"Cipher"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientIpPool         string `position:"Query" name:"ClientIpPool"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Compress             string `position:"Query" name:"Compress"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	Port                 int64  `position:"Query" name:"Port"`
	Proto                string `position:"Query" name:"Proto"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateSslVpnServerRequest) Invoke(client goaliyun.Client) (*CreateSslVpnServerResponse, error) {
	resp := &CreateSslVpnServerResponse{}
	err := client.Request("vpc", "CreateSslVpnServer", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSslVpnServerResponse struct {
	RequestId      goaliyun.String
	SslVpnServerId goaliyun.String
	Name           goaliyun.String
}
