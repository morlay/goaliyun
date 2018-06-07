package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVirtualBorderRoutersRequest struct {
	Filters              *DescribeVirtualBorderRoutersFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                   `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                  `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int64                                   `position:"Query" name:"PageSize"`
	OwnerId              int64                                   `position:"Query" name:"OwnerId"`
	PageNumber           int64                                   `position:"Query" name:"PageNumber"`
	RegionId             string                                  `position:"Query" name:"RegionId"`
}

func (req *DescribeVirtualBorderRoutersRequest) Invoke(client goaliyun.Client) (*DescribeVirtualBorderRoutersResponse, error) {
	resp := &DescribeVirtualBorderRoutersResponse{}
	err := client.Request("ecs", "DescribeVirtualBorderRouters", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVirtualBorderRoutersFilter struct {
	Key    string                                 `name:"Key"`
	Values *DescribeVirtualBorderRoutersValueList `type:"Repeated" name:"Value"`
}

type DescribeVirtualBorderRoutersResponse struct {
	RequestId              goaliyun.String
	PageNumber             goaliyun.Integer
	PageSize               goaliyun.Integer
	TotalCount             goaliyun.Integer
	VirtualBorderRouterSet DescribeVirtualBorderRoutersVirtualBorderRouterTypeList
}

type DescribeVirtualBorderRoutersVirtualBorderRouterType struct {
	VbrId                            goaliyun.String
	CreationTime                     goaliyun.String
	ActivationTime                   goaliyun.String
	TerminationTime                  goaliyun.String
	RecoveryTime                     goaliyun.String
	Status                           goaliyun.String
	VlanId                           goaliyun.Integer
	CircuitCode                      goaliyun.String
	RouteTableId                     goaliyun.String
	VlanInterfaceId                  goaliyun.String
	LocalGatewayIp                   goaliyun.String
	PeerGatewayIp                    goaliyun.String
	PeeringSubnetMask                goaliyun.String
	PhysicalConnectionId             goaliyun.String
	PhysicalConnectionStatus         goaliyun.String
	PhysicalConnectionBusinessStatus goaliyun.String
	PhysicalConnectionOwnerUid       goaliyun.String
	AccessPointId                    goaliyun.String
	Name                             goaliyun.String
	Description                      goaliyun.String
}

type DescribeVirtualBorderRoutersFilterList []DescribeVirtualBorderRoutersFilter

func (list *DescribeVirtualBorderRoutersFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersFilter)
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

type DescribeVirtualBorderRoutersValueList []string

func (list *DescribeVirtualBorderRoutersValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeVirtualBorderRoutersVirtualBorderRouterTypeList []DescribeVirtualBorderRoutersVirtualBorderRouterType

func (list *DescribeVirtualBorderRoutersVirtualBorderRouterTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersVirtualBorderRouterType)
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
