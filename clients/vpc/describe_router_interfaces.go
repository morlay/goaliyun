package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRouterInterfacesRequest struct {
	Filters              *DescribeRouterInterfacesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                               `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int64                               `position:"Query" name:"PageSize"`
	OwnerId              int64                               `position:"Query" name:"OwnerId"`
	PageNumber           int64                               `position:"Query" name:"PageNumber"`
	RegionId             string                              `position:"Query" name:"RegionId"`
}

func (req *DescribeRouterInterfacesRequest) Invoke(client goaliyun.Client) (*DescribeRouterInterfacesResponse, error) {
	resp := &DescribeRouterInterfacesResponse{}
	err := client.Request("vpc", "DescribeRouterInterfaces", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRouterInterfacesFilter struct {
	Key    string                             `name:"Key"`
	Values *DescribeRouterInterfacesValueList `type:"Repeated" name:"Value"`
}

type DescribeRouterInterfacesResponse struct {
	RequestId          goaliyun.String
	PageNumber         goaliyun.Integer
	PageSize           goaliyun.Integer
	TotalCount         goaliyun.Integer
	RouterInterfaceSet DescribeRouterInterfacesRouterInterfaceTypeList
}

type DescribeRouterInterfacesRouterInterfaceType struct {
	RouterInterfaceId               goaliyun.String
	OppositeRegionId                goaliyun.String
	Role                            goaliyun.String
	Spec                            goaliyun.String
	Name                            goaliyun.String
	Description                     goaliyun.String
	RouterId                        goaliyun.String
	RouterType                      goaliyun.String
	CreationTime                    goaliyun.String
	EndTime                         goaliyun.String
	ChargeType                      goaliyun.String
	Status                          goaliyun.String
	BusinessStatus                  goaliyun.String
	ConnectedTime                   goaliyun.String
	OppositeInterfaceId             goaliyun.String
	OppositeInterfaceSpec           goaliyun.String
	OppositeInterfaceStatus         goaliyun.String
	OppositeInterfaceBusinessStatus goaliyun.String
	OppositeRouterId                goaliyun.String
	OppositeRouterType              goaliyun.String
	OppositeInterfaceOwnerId        goaliyun.String
	AccessPointId                   goaliyun.String
	OppositeAccessPointId           goaliyun.String
	HealthCheckSourceIp             goaliyun.String
	HealthCheckTargetIp             goaliyun.String
	OppositeVpcInstanceId           goaliyun.String
	VpcInstanceId                   goaliyun.String
}

type DescribeRouterInterfacesFilterList []DescribeRouterInterfacesFilter

func (list *DescribeRouterInterfacesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouterInterfacesFilter)
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

type DescribeRouterInterfacesValueList []string

func (list *DescribeRouterInterfacesValueList) UnmarshalJSON(data []byte) error {
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

type DescribeRouterInterfacesRouterInterfaceTypeList []DescribeRouterInterfacesRouterInterfaceType

func (list *DescribeRouterInterfacesRouterInterfaceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouterInterfacesRouterInterfaceType)
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
