package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSslVpnServersRequest struct {
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSslVpnServersRequest) Invoke(client goaliyun.Client) (*DescribeSslVpnServersResponse, error) {
	resp := &DescribeSslVpnServersResponse{}
	err := client.Request("vpc", "DescribeSslVpnServers", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSslVpnServersResponse struct {
	RequestId     goaliyun.String
	TotalCount    goaliyun.Integer
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	SslVpnServers DescribeSslVpnServersSslVpnServerList
}

type DescribeSslVpnServersSslVpnServer struct {
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

type DescribeSslVpnServersSslVpnServerList []DescribeSslVpnServersSslVpnServer

func (list *DescribeSslVpnServersSslVpnServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSslVpnServersSslVpnServer)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
