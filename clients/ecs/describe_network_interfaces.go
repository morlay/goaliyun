package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeNetworkInterfacesRequest struct {
	ResourceOwnerId      int64                                            `position:"Query" name:"ResourceOwnerId"`
	SecurityGroupId      string                                           `position:"Query" name:"SecurityGroupId"`
	Type                 string                                           `position:"Query" name:"Type"`
	PageNumber           int64                                            `position:"Query" name:"PageNumber"`
	PageSize             int64                                            `position:"Query" name:"PageSize"`
	NetworkInterfaceName string                                           `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string                                           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                           `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                            `position:"Query" name:"OwnerId"`
	VSwitchId            string                                           `position:"Query" name:"VSwitchId"`
	InstanceId           string                                           `position:"Query" name:"InstanceId"`
	PrimaryIpAddress     string                                           `position:"Query" name:"PrimaryIpAddress"`
	NetworkInterfaceIds  *DescribeNetworkInterfacesNetworkInterfaceIdList `position:"Query" type:"Repeated" name:"NetworkInterfaceId"`
	RegionId             string                                           `position:"Query" name:"RegionId"`
}

func (req *DescribeNetworkInterfacesRequest) Invoke(client goaliyun.Client) (*DescribeNetworkInterfacesResponse, error) {
	resp := &DescribeNetworkInterfacesResponse{}
	err := client.Request("ecs", "DescribeNetworkInterfaces", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNetworkInterfacesResponse struct {
	RequestId            goaliyun.String
	TotalCount           goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	NetworkInterfaceSets DescribeNetworkInterfacesNetworkInterfaceSetList
}

type DescribeNetworkInterfacesNetworkInterfaceSet struct {
	NetworkInterfaceId   goaliyun.String
	Status               goaliyun.String
	Type                 goaliyun.String
	VpcId                goaliyun.String
	VSwitchId            goaliyun.String
	ZoneId               goaliyun.String
	PrivateIpAddress     goaliyun.String
	MacAddress           goaliyun.String
	NetworkInterfaceName goaliyun.String
	Description          goaliyun.String
	InstanceId           goaliyun.String
	CreationTime         goaliyun.String
	PrivateIpSets        DescribeNetworkInterfacesPrivateIpSetList
	SecurityGroupIds     DescribeNetworkInterfacesSecurityGroupIdList
	AssociatedPublicIp   DescribeNetworkInterfacesAssociatedPublicIp
}

type DescribeNetworkInterfacesPrivateIpSet struct {
	PrivateIpAddress    goaliyun.String
	Primary             bool
	AssociatedPublicIp1 DescribeNetworkInterfacesAssociatedPublicIp1
}

type DescribeNetworkInterfacesAssociatedPublicIp1 struct {
	PublicIpAddress goaliyun.String
	AllocationId    goaliyun.String
}

type DescribeNetworkInterfacesAssociatedPublicIp struct {
	PublicIpAddress goaliyun.String
	AllocationId    goaliyun.String
}

type DescribeNetworkInterfacesNetworkInterfaceIdList []string

func (list *DescribeNetworkInterfacesNetworkInterfaceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNetworkInterfacesNetworkInterfaceSetList []DescribeNetworkInterfacesNetworkInterfaceSet

func (list *DescribeNetworkInterfacesNetworkInterfaceSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesNetworkInterfaceSet)
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

type DescribeNetworkInterfacesPrivateIpSetList []DescribeNetworkInterfacesPrivateIpSet

func (list *DescribeNetworkInterfacesPrivateIpSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesPrivateIpSet)
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

type DescribeNetworkInterfacesSecurityGroupIdList []goaliyun.String

func (list *DescribeNetworkInterfacesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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
