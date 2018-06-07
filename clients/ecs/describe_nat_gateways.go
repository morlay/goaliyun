package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeNatGatewaysRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeNatGatewaysRequest) Invoke(client goaliyun.Client) (*DescribeNatGatewaysResponse, error) {
	resp := &DescribeNatGatewaysResponse{}
	err := client.Request("ecs", "DescribeNatGateways", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNatGatewaysResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	NatGateways DescribeNatGatewaysNatGatewayList
}

type DescribeNatGatewaysNatGateway struct {
	NatGatewayId        goaliyun.String
	RegionId            goaliyun.String
	Name                goaliyun.String
	Description         goaliyun.String
	VpcId               goaliyun.String
	Spec                goaliyun.String
	InstanceChargeType  goaliyun.String
	BusinessStatus      goaliyun.String
	CreationTime        goaliyun.String
	Status              goaliyun.String
	ForwardTableIds     DescribeNatGatewaysForwardTableIdList
	BandwidthPackageIds DescribeNatGatewaysBandwidthPackageIdList
}

type DescribeNatGatewaysNatGatewayList []DescribeNatGatewaysNatGateway

func (list *DescribeNatGatewaysNatGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNatGatewaysNatGateway)
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

type DescribeNatGatewaysForwardTableIdList []goaliyun.String

func (list *DescribeNatGatewaysForwardTableIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type DescribeNatGatewaysBandwidthPackageIdList []goaliyun.String

func (list *DescribeNatGatewaysBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
