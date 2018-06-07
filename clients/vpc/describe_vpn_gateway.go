package vpc

import (
	"github.com/morlay/goaliyun"
)

type DescribeVpnGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpnGatewayRequest) Invoke(client goaliyun.Client) (*DescribeVpnGatewayResponse, error) {
	resp := &DescribeVpnGatewayResponse{}
	err := client.Request("vpc", "DescribeVpnGateway", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpnGatewayResponse struct {
	RequestId         goaliyun.String
	VpnGatewayId      goaliyun.String
	VpcId             goaliyun.String
	VSwitchId         goaliyun.String
	InternetIp        goaliyun.String
	CreateTime        goaliyun.Integer
	EndTime           goaliyun.Integer
	Spec              goaliyun.String
	Name              goaliyun.String
	Description       goaliyun.String
	Status            goaliyun.String
	BusinessStatus    goaliyun.String
	ChargeType        goaliyun.String
	IpsecVpn          goaliyun.String
	SslVpn            goaliyun.String
	SslMaxConnections goaliyun.Integer
}
