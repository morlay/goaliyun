package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVirtualBorderRoutersForPhysicalConnectionRequest struct {
	Filters              *DescribeVirtualBorderRoutersForPhysicalConnectionFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                                        `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                                       `position:"Query" name:"ResourceOwnerAccount"`
	PhysicalConnectionId string                                                       `position:"Query" name:"PhysicalConnectionId"`
	PageSize             int64                                                        `position:"Query" name:"PageSize"`
	OwnerId              int64                                                        `position:"Query" name:"OwnerId"`
	PageNumber           int64                                                        `position:"Query" name:"PageNumber"`
	RegionId             string                                                       `position:"Query" name:"RegionId"`
}

func (req *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) Invoke(client goaliyun.Client) (*DescribeVirtualBorderRoutersForPhysicalConnectionResponse, error) {
	resp := &DescribeVirtualBorderRoutersForPhysicalConnectionResponse{}
	err := client.Request("vpc", "DescribeVirtualBorderRoutersForPhysicalConnection", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVirtualBorderRoutersForPhysicalConnectionFilter struct {
	Key    string                                                      `name:"Key"`
	Values *DescribeVirtualBorderRoutersForPhysicalConnectionValueList `type:"Repeated" name:"Value"`
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponse struct {
	RequestId                                   goaliyun.String
	PageNumber                                  goaliyun.Integer
	PageSize                                    goaliyun.Integer
	TotalCount                                  goaliyun.Integer
	VirtualBorderRouterForPhysicalConnectionSet DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionTypeList
}

type DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionType struct {
	VbrId           goaliyun.String
	VbrOwnerUid     goaliyun.Integer
	CreationTime    goaliyun.String
	ActivationTime  goaliyun.String
	TerminationTime goaliyun.String
	RecoveryTime    goaliyun.String
	VlanId          goaliyun.Integer
	CircuitCode     goaliyun.String
}

type DescribeVirtualBorderRoutersForPhysicalConnectionFilterList []DescribeVirtualBorderRoutersForPhysicalConnectionFilter

func (list *DescribeVirtualBorderRoutersForPhysicalConnectionFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersForPhysicalConnectionFilter)
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

type DescribeVirtualBorderRoutersForPhysicalConnectionValueList []string

func (list *DescribeVirtualBorderRoutersForPhysicalConnectionValueList) UnmarshalJSON(data []byte) error {
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

type DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionTypeList []DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionType

func (list *DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionType)
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
