package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLaunchTemplateVersionsRequest struct {
	LaunchTemplateName     string                                                   `position:"Query" name:"LaunchTemplateName"`
	MaxVersion             int64                                                    `position:"Query" name:"MaxVersion"`
	ResourceOwnerId        int64                                                    `position:"Query" name:"ResourceOwnerId"`
	DefaultVersion         string                                                   `position:"Query" name:"DefaultVersion"`
	MinVersion             int64                                                    `position:"Query" name:"MinVersion"`
	PageNumber             int64                                                    `position:"Query" name:"PageNumber"`
	PageSize               int64                                                    `position:"Query" name:"PageSize"`
	LaunchTemplateId       string                                                   `position:"Query" name:"LaunchTemplateId"`
	ResourceOwnerAccount   string                                                   `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string                                                   `position:"Query" name:"OwnerAccount"`
	OwnerId                int64                                                    `position:"Query" name:"OwnerId"`
	LaunchTemplateVersions *DescribeLaunchTemplateVersionsLaunchTemplateVersionList `position:"Query" type:"Repeated" name:"LaunchTemplateVersion"`
	DetailFlag             string                                                   `position:"Query" name:"DetailFlag"`
	RegionId               string                                                   `position:"Query" name:"RegionId"`
}

func (req *DescribeLaunchTemplateVersionsRequest) Invoke(client goaliyun.Client) (*DescribeLaunchTemplateVersionsResponse, error) {
	resp := &DescribeLaunchTemplateVersionsResponse{}
	err := client.Request("ecs", "DescribeLaunchTemplateVersions", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLaunchTemplateVersionsResponse struct {
	RequestId                 goaliyun.String
	TotalCount                goaliyun.Integer
	PageNumber                goaliyun.Integer
	PageSize                  goaliyun.Integer
	LaunchTemplateVersionSets DescribeLaunchTemplateVersionsLaunchTemplateVersionSetList
}

type DescribeLaunchTemplateVersionsLaunchTemplateVersionSet struct {
	CreateTime         goaliyun.String
	ModifiedTime       goaliyun.String
	LaunchTemplateId   goaliyun.String
	LaunchTemplateName goaliyun.String
	DefaultVersion     bool
	VersionNumber      goaliyun.Integer
	VersionDescription goaliyun.String
	CreatedBy          goaliyun.String
	LaunchTemplateData DescribeLaunchTemplateVersionsLaunchTemplateData
}

type DescribeLaunchTemplateVersionsLaunchTemplateData struct {
	ImageId                     goaliyun.String
	ImageOwnerAlias             goaliyun.String
	InstanceType                goaliyun.String
	SecurityGroupId             goaliyun.String
	VpcId                       goaliyun.String
	VSwitchId                   goaliyun.String
	InstanceName                goaliyun.String
	Description                 goaliyun.String
	InternetMaxBandwidthIn      goaliyun.Integer
	InternetMaxBandwidthOut     goaliyun.Integer
	HostName                    goaliyun.String
	ZoneId                      goaliyun.String
	SystemDiskSize              goaliyun.Integer
	SystemDiskCategory          goaliyun.String
	SystemDiskDiskName          goaliyun.String
	SystemDiskDescription       goaliyun.String
	SystemDiskIops              goaliyun.Integer
	IoOptimized                 goaliyun.String
	InstanceChargeType          goaliyun.String
	Period                      goaliyun.Integer
	InternetChargeType          goaliyun.String
	EnableVmOsConfig            bool
	NetworkType                 goaliyun.String
	UserData                    goaliyun.String
	KeyPairName                 goaliyun.String
	RamRoleName                 goaliyun.String
	AutoReleaseTime             goaliyun.String
	SpotStrategy                goaliyun.String
	SpotPriceLimit              goaliyun.Float
	SpotDuration                goaliyun.Integer
	ResourceGroupId             goaliyun.String
	SecurityEnhancementStrategy bool
	DataDisks                   DescribeLaunchTemplateVersionsDataDiskList
	NetworkInterfaces           DescribeLaunchTemplateVersionsNetworkInterfaceList
	Tags                        DescribeLaunchTemplateVersionsInstanceTagList
}

type DescribeLaunchTemplateVersionsDataDisk struct {
	Size               goaliyun.Integer
	SnapshotId         goaliyun.String
	Category           goaliyun.String
	Encrypted          goaliyun.String
	DiskName           goaliyun.String
	Description        goaliyun.String
	DeleteWithInstance bool
}

type DescribeLaunchTemplateVersionsNetworkInterface struct {
	PrimaryIpAddress     goaliyun.String
	VSwitchId            goaliyun.String
	SecurityGroupId      goaliyun.String
	NetworkInterfaceName goaliyun.String
	Description          goaliyun.String
}

type DescribeLaunchTemplateVersionsInstanceTag struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type DescribeLaunchTemplateVersionsLaunchTemplateVersionList []int64

func (list *DescribeLaunchTemplateVersionsLaunchTemplateVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type DescribeLaunchTemplateVersionsLaunchTemplateVersionSetList []DescribeLaunchTemplateVersionsLaunchTemplateVersionSet

func (list *DescribeLaunchTemplateVersionsLaunchTemplateVersionSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLaunchTemplateVersionsLaunchTemplateVersionSet)
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

type DescribeLaunchTemplateVersionsDataDiskList []DescribeLaunchTemplateVersionsDataDisk

func (list *DescribeLaunchTemplateVersionsDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLaunchTemplateVersionsDataDisk)
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

type DescribeLaunchTemplateVersionsNetworkInterfaceList []DescribeLaunchTemplateVersionsNetworkInterface

func (list *DescribeLaunchTemplateVersionsNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLaunchTemplateVersionsNetworkInterface)
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

type DescribeLaunchTemplateVersionsInstanceTagList []DescribeLaunchTemplateVersionsInstanceTag

func (list *DescribeLaunchTemplateVersionsInstanceTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLaunchTemplateVersionsInstanceTag)
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
