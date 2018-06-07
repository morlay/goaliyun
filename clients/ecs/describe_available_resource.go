package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAvailableResourceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	IoOptimized          string `position:"Query" name:"IoOptimized"`
	DataDiskCategory     string `position:"Query" name:"DataDiskCategory"`
	SystemDiskCategory   string `position:"Query" name:"SystemDiskCategory"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	NetworkCategory      string `position:"Query" name:"NetworkCategory"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DedicatedHostId      string `position:"Query" name:"DedicatedHostId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpotStrategy         string `position:"Query" name:"SpotStrategy"`
	DestinationResource  string `position:"Query" name:"DestinationResource"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAvailableResourceRequest) Invoke(client goaliyun.Client) (*DescribeAvailableResourceResponse, error) {
	resp := &DescribeAvailableResourceResponse{}
	err := client.Request("ecs", "DescribeAvailableResource", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAvailableResourceResponse struct {
	RequestId      goaliyun.String
	AvailableZones DescribeAvailableResourceAvailableZoneList
}

type DescribeAvailableResourceAvailableZone struct {
	RegionId           goaliyun.String
	ZoneId             goaliyun.String
	Status             goaliyun.String
	AvailableResources DescribeAvailableResourceAvailableResourceList
}

type DescribeAvailableResourceAvailableResource struct {
	Type               goaliyun.String
	SupportedResources DescribeAvailableResourceSupportedResourceList
}

type DescribeAvailableResourceSupportedResource struct {
	Value  goaliyun.String
	Status goaliyun.String
	Min    goaliyun.Integer
	Max    goaliyun.Integer
	Unit   goaliyun.String
}

type DescribeAvailableResourceAvailableZoneList []DescribeAvailableResourceAvailableZone

func (list *DescribeAvailableResourceAvailableZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableResourceAvailableZone)
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

type DescribeAvailableResourceAvailableResourceList []DescribeAvailableResourceAvailableResource

func (list *DescribeAvailableResourceAvailableResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableResourceAvailableResource)
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

type DescribeAvailableResourceSupportedResourceList []DescribeAvailableResourceSupportedResource

func (list *DescribeAvailableResourceSupportedResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableResourceSupportedResource)
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
