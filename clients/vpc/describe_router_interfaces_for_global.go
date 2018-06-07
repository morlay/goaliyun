package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRouterInterfacesForGlobalRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRouterInterfacesForGlobalRequest) Invoke(client goaliyun.Client) (*DescribeRouterInterfacesForGlobalResponse, error) {
	resp := &DescribeRouterInterfacesForGlobalResponse{}
	err := client.Request("vpc", "DescribeRouterInterfacesForGlobal", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRouterInterfacesForGlobalResponse struct {
	RequestId          goaliyun.String
	Code               goaliyun.String
	Message            goaliyun.String
	Desc               goaliyun.String
	Success            bool
	PageSize           goaliyun.Integer
	PageNumber         goaliyun.Integer
	TotalCount         goaliyun.Integer
	RouterInterfaceSet DescribeRouterInterfacesForGlobalRouterInterfaceTypeList
}

type DescribeRouterInterfacesForGlobalRouterInterfaceType struct {
	BusinessStatus                  goaliyun.String
	AccessPointId                   goaliyun.String
	ChargeType                      goaliyun.String
	ConnectedTime                   goaliyun.String
	CreationTime                    goaliyun.String
	RouterInterfaceId               goaliyun.String
	OppositeInterfaceBusinessStatus goaliyun.String
	OppositeInterfaceId             goaliyun.String
	OppositeInterfaceOwnerId        goaliyun.Integer
	OppositeInterfaceSpec           goaliyun.String
	OppositeInterfaceStatus         goaliyun.String
	OppositeRegionId                goaliyun.String
	OppositeAccessPointId           goaliyun.String
	OppositeRouterId                goaliyun.String
	OppositeRouterType              goaliyun.String
	OppositeVpcInstanceId           goaliyun.String
	RegionId                        goaliyun.String
	Role                            goaliyun.String
	RouterId                        goaliyun.String
	RouterType                      goaliyun.String
	Spec                            goaliyun.String
	Status                          goaliyun.String
	VpcInstanceId                   goaliyun.String
	Name                            goaliyun.String
	Description                     goaliyun.String
	HealthCheckSourceIp             goaliyun.String
	HealthCheckTargetIp             goaliyun.String
}

type DescribeRouterInterfacesForGlobalRouterInterfaceTypeList []DescribeRouterInterfacesForGlobalRouterInterfaceType

func (list *DescribeRouterInterfacesForGlobalRouterInterfaceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouterInterfacesForGlobalRouterInterfaceType)
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
