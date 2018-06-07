package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeResourcesModificationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateAcrossZone    string `position:"Query" name:"MigrateAcrossZone"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OperationType        string `position:"Query" name:"OperationType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DestinationResource  string `position:"Query" name:"DestinationResource"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeResourcesModificationRequest) Invoke(client goaliyun.Client) (*DescribeResourcesModificationResponse, error) {
	resp := &DescribeResourcesModificationResponse{}
	err := client.Request("ecs", "DescribeResourcesModification", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeResourcesModificationResponse struct {
	RequestId      goaliyun.String
	AvailableZones DescribeResourcesModificationAvailableZoneList
}

type DescribeResourcesModificationAvailableZone struct {
	RegionId           goaliyun.String
	ZoneId             goaliyun.String
	Status             goaliyun.String
	AvailableResources DescribeResourcesModificationAvailableResourceList
}

type DescribeResourcesModificationAvailableResource struct {
	Type               goaliyun.String
	SupportedResources DescribeResourcesModificationSupportedResourceList
}

type DescribeResourcesModificationSupportedResource struct {
	Value  goaliyun.String
	Status goaliyun.String
	Min    goaliyun.Integer
	Max    goaliyun.Integer
	Unit   goaliyun.String
}

type DescribeResourcesModificationAvailableZoneList []DescribeResourcesModificationAvailableZone

func (list *DescribeResourcesModificationAvailableZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourcesModificationAvailableZone)
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

type DescribeResourcesModificationAvailableResourceList []DescribeResourcesModificationAvailableResource

func (list *DescribeResourcesModificationAvailableResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourcesModificationAvailableResource)
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

type DescribeResourcesModificationSupportedResourceList []DescribeResourcesModificationSupportedResource

func (list *DescribeResourcesModificationSupportedResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourcesModificationSupportedResource)
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
