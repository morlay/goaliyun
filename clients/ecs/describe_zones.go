package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeZonesRequest struct {
	SpotStrategy         string `position:"Query" name:"SpotStrategy"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	Verbose              string `position:"Query" name:"Verbose"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeZonesRequest) Invoke(client goaliyun.Client) (*DescribeZonesResponse, error) {
	resp := &DescribeZonesResponse{}
	err := client.Request("ecs", "DescribeZones", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeZonesResponse struct {
	RequestId goaliyun.String
	Zones     DescribeZonesZoneList
}

type DescribeZonesZone struct {
	ZoneId                      goaliyun.String
	LocalName                   goaliyun.String
	AvailableResources          DescribeZonesResourcesInfoList
	AvailableResourceCreation   DescribeZonesAvailableResourceCreationList
	AvailableDiskCategories     DescribeZonesAvailableDiskCategoryList
	AvailableInstanceTypes      DescribeZonesAvailableInstanceTypeList
	AvailableVolumeCategories   DescribeZonesAvailableVolumeCategoryList
	AvailableDedicatedHostTypes DescribeZonesAvailableDedicatedHostTypeList
	DedicatedHostGenerations    DescribeZonesDedicatedHostGenerationList
}

type DescribeZonesResourcesInfo struct {
	IoOptimized          bool
	SystemDiskCategories DescribeZonesSystemDiskCategoryList
	DataDiskCategories   DescribeZonesDataDiskCategoryList
	NetworkTypes         DescribeZonesNetworkTypeList
	InstanceTypes        DescribeZonesInstanceTypeList
	InstanceTypeFamilies DescribeZonesInstanceTypeFamilyList
	InstanceGenerations  DescribeZonesInstanceGenerationList
}

type DescribeZonesZoneList []DescribeZonesZone

func (list *DescribeZonesZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesZone)
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

type DescribeZonesResourcesInfoList []DescribeZonesResourcesInfo

func (list *DescribeZonesResourcesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesResourcesInfo)
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

type DescribeZonesAvailableResourceCreationList []goaliyun.String

func (list *DescribeZonesAvailableResourceCreationList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableDiskCategoryList []goaliyun.String

func (list *DescribeZonesAvailableDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableInstanceTypeList []goaliyun.String

func (list *DescribeZonesAvailableInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableVolumeCategoryList []goaliyun.String

func (list *DescribeZonesAvailableVolumeCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableDedicatedHostTypeList []goaliyun.String

func (list *DescribeZonesAvailableDedicatedHostTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesDedicatedHostGenerationList []goaliyun.String

func (list *DescribeZonesDedicatedHostGenerationList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesSystemDiskCategoryList []goaliyun.String

func (list *DescribeZonesSystemDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesDataDiskCategoryList []goaliyun.String

func (list *DescribeZonesDataDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesNetworkTypeList []goaliyun.String

func (list *DescribeZonesNetworkTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceTypeList []goaliyun.String

func (list *DescribeZonesInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceTypeFamilyList []goaliyun.String

func (list *DescribeZonesInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceGenerationList []goaliyun.String

func (list *DescribeZonesInstanceGenerationList) UnmarshalJSON(data []byte) error {
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
