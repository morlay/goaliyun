package vpc

import (
	"github.com/morlay/goaliyun"
)

type DownloadVpnConnectionConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DownloadVpnConnectionConfigRequest) Invoke(client goaliyun.Client) (*DownloadVpnConnectionConfigResponse, error) {
	resp := &DownloadVpnConnectionConfigResponse{}
	err := client.Request("vpc", "DownloadVpnConnectionConfig", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DownloadVpnConnectionConfigResponse struct {
	RequestId           goaliyun.String
	VpnConnectionConfig DownloadVpnConnectionConfigVpnConnectionConfig
}

type DownloadVpnConnectionConfigVpnConnectionConfig struct {
	LocalSubnet  goaliyun.String
	RemoteSubnet goaliyun.String
	Local        goaliyun.String
	Remote       goaliyun.String
	IkeConfig    DownloadVpnConnectionConfigIkeConfig
	IpsecConfig  DownloadVpnConnectionConfigIpsecConfig
}

type DownloadVpnConnectionConfigIkeConfig struct {
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

type DownloadVpnConnectionConfigIpsecConfig struct {
	IpsecEncAlg   goaliyun.String
	IpsecAuthAlg  goaliyun.String
	IpsecPfs      goaliyun.String
	IpsecLifetime goaliyun.Integer
}
