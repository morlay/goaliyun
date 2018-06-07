package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVpnGatewaysRequest struct {
	BusinessStatus       string `position:"Query" name:"BusinessStatus"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVpnGatewaysRequest) Invoke(client goaliyun.Client) (*DescribeVpnGatewaysResponse, error) {
	resp := &DescribeVpnGatewaysResponse{}
	err := client.Request("vpc", "DescribeVpnGateways", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVpnGatewaysResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	VpnGateways DescribeVpnGatewaysVpnGatewayList
}

type DescribeVpnGatewaysVpnGateway struct {
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

type DescribeVpnGatewaysVpnGatewayList []DescribeVpnGatewaysVpnGateway

func (list *DescribeVpnGatewaysVpnGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpnGatewaysVpnGateway)
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
