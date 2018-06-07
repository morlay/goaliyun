package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVpnConnectionsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpnConnectionsRequest) Invoke(client goaliyun.Client) (*DescribeVpnConnectionsResponse, error) {
	resp := &DescribeVpnConnectionsResponse{}
	err := client.Request("vpc", "DescribeVpnConnections", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpnConnectionsResponse struct {
	RequestId      goaliyun.String
	TotalCount     goaliyun.Integer
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	VpnConnections DescribeVpnConnectionsVpnConnectionList
}

type DescribeVpnConnectionsVpnConnection struct {
	VpnConnectionId   goaliyun.String
	CustomerGatewayId goaliyun.String
	VpnGatewayId      goaliyun.String
	Name              goaliyun.String
	LocalSubnet       goaliyun.String
	RemoteSubnet      goaliyun.String
	CreateTime        goaliyun.Integer
	EffectImmediately bool
	Status            goaliyun.String
	IkeConfig         DescribeVpnConnectionsIkeConfig
	IpsecConfig       DescribeVpnConnectionsIpsecConfig
}

type DescribeVpnConnectionsIkeConfig struct {
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

type DescribeVpnConnectionsIpsecConfig struct {
	IpsecEncAlg   goaliyun.String
	IpsecAuthAlg  goaliyun.String
	IpsecPfs      goaliyun.String
	IpsecLifetime goaliyun.Integer
}

type DescribeVpnConnectionsVpnConnectionList []DescribeVpnConnectionsVpnConnection

func (list *DescribeVpnConnectionsVpnConnectionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpnConnectionsVpnConnection)
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
