package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePhysicalConnectionsRequest struct {
	Filters              *DescribePhysicalConnectionsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                 `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                 `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                 `position:"Query" name:"OwnerAccount"`
	PageSize             int64                                  `position:"Query" name:"PageSize"`
	OwnerId              int64                                  `position:"Query" name:"OwnerId"`
	PageNumber           int64                                  `position:"Query" name:"PageNumber"`
	RegionId             string                                 `position:"Query" name:"RegionId"`
}

func (req *DescribePhysicalConnectionsRequest) Invoke(client goaliyun.Client) (*DescribePhysicalConnectionsResponse, error) {
	resp := &DescribePhysicalConnectionsResponse{}
	err := client.Request("vpc", "DescribePhysicalConnections", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePhysicalConnectionsFilter struct {
	Key    string                                `name:"Key"`
	Values *DescribePhysicalConnectionsValueList `type:"Repeated" name:"Value"`
}

type DescribePhysicalConnectionsResponse struct {
	RequestId             goaliyun.String
	PageNumber            goaliyun.Integer
	PageSize              goaliyun.Integer
	TotalCount            goaliyun.Integer
	PhysicalConnectionSet DescribePhysicalConnectionsPhysicalConnectionTypeList
}

type DescribePhysicalConnectionsPhysicalConnectionType struct {
	PhysicalConnectionId          goaliyun.String
	AccessPointId                 goaliyun.String
	Type                          goaliyun.String
	Status                        goaliyun.String
	BusinessStatus                goaliyun.String
	CreationTime                  goaliyun.String
	EnabledTime                   goaliyun.String
	LineOperator                  goaliyun.String
	Spec                          goaliyun.String
	PeerLocation                  goaliyun.String
	PortType                      goaliyun.String
	RedundantPhysicalConnectionId goaliyun.String
	Name                          goaliyun.String
	Description                   goaliyun.String
	AdLocation                    goaliyun.String
	PortNumber                    goaliyun.String
	CircuitCode                   goaliyun.String
	Bandwidth                     goaliyun.Integer
}

type DescribePhysicalConnectionsFilterList []DescribePhysicalConnectionsFilter

func (list *DescribePhysicalConnectionsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePhysicalConnectionsFilter)
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

type DescribePhysicalConnectionsValueList []string

func (list *DescribePhysicalConnectionsValueList) UnmarshalJSON(data []byte) error {
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

type DescribePhysicalConnectionsPhysicalConnectionTypeList []DescribePhysicalConnectionsPhysicalConnectionType

func (list *DescribePhysicalConnectionsPhysicalConnectionTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePhysicalConnectionsPhysicalConnectionType)
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
