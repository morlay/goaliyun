package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RunInstancesRequest struct {
	LaunchTemplateName          string                            `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId             int64                             `position:"Query" name:"ResourceOwnerId"`
	HpcClusterId                string                            `position:"Query" name:"HpcClusterId"`
	SecurityEnhancementStrategy string                            `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                            `position:"Query" name:"KeyPairName"`
	SpotPriceLimit              float64                           `position:"Query" name:"SpotPriceLimit"`
	HostName                    string                            `position:"Query" name:"HostName"`
	Password                    string                            `position:"Query" name:"Password"`
	Tags                        *RunInstancesTagList              `position:"Query" type:"Repeated" name:"Tag"`
	DryRun                      string                            `position:"Query" name:"DryRun"`
	LaunchTemplateId            string                            `position:"Query" name:"LaunchTemplateId"`
	OwnerId                     int64                             `position:"Query" name:"OwnerId"`
	VSwitchId                   string                            `position:"Query" name:"VSwitchId"`
	SpotStrategy                string                            `position:"Query" name:"SpotStrategy"`
	InstanceName                string                            `position:"Query" name:"InstanceName"`
	InternetChargeType          string                            `position:"Query" name:"InternetChargeType"`
	ZoneId                      string                            `position:"Query" name:"ZoneId"`
	InternetMaxBandwidthIn      int64                             `position:"Query" name:"InternetMaxBandwidthIn"`
	ImageId                     string                            `position:"Query" name:"ImageId"`
	SpotInterruptionBehavior    string                            `position:"Query" name:"SpotInterruptionBehavior"`
	ClientToken                 string                            `position:"Query" name:"ClientToken"`
	IoOptimized                 string                            `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                            `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     int64                             `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                 string                            `position:"Query" name:"Description"`
	SystemDiskCategory          string                            `position:"Query" name:"SystemDiskCategory"`
	UserData                    string                            `position:"Query" name:"UserData"`
	InstanceType                string                            `position:"Query" name:"InstanceType"`
	NetworkInterfaces           *RunInstancesNetworkInterfaceList `position:"Query" type:"Repeated" name:"NetworkInterface"`
	Amount                      int64                             `position:"Query" name:"Amount"`
	ResourceOwnerAccount        string                            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                            `position:"Query" name:"OwnerAccount"`
	SystemDiskDiskName          string                            `position:"Query" name:"SystemDiskDiskName"`
	RamRoleName                 string                            `position:"Query" name:"RamRoleName"`
	AutoReleaseTime             string                            `position:"Query" name:"AutoReleaseTime"`
	DedicatedHostId             string                            `position:"Query" name:"DedicatedHostId"`
	DataDisks                   *RunInstancesDataDiskList         `position:"Query" type:"Repeated" name:"DataDisk"`
	LaunchTemplateVersion       int64                             `position:"Query" name:"LaunchTemplateVersion"`
	SystemDiskSize              string                            `position:"Query" name:"SystemDiskSize"`
	SystemDiskDescription       string                            `position:"Query" name:"SystemDiskDescription"`
	RegionId                    string                            `position:"Query" name:"RegionId"`
}

func (req *RunInstancesRequest) Invoke(client goaliyun.Client) (*RunInstancesResponse, error) {
	resp := &RunInstancesResponse{}
	err := client.Request("ecs", "RunInstances", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RunInstancesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type RunInstancesNetworkInterface struct {
	PrimaryIpAddress     string `name:"PrimaryIpAddress"`
	VSwitchId            string `name:"VSwitchId"`
	SecurityGroupId      string `name:"SecurityGroupId"`
	NetworkInterfaceName string `name:"NetworkInterfaceName"`
	Description          string `name:"Description"`
}

type RunInstancesDataDisk struct {
	Size               int64  `name:"Size"`
	SnapshotId         string `name:"SnapshotId"`
	Category           string `name:"Category"`
	Encrypted          string `name:"Encrypted"`
	DiskName           string `name:"DiskName"`
	Description        string `name:"Description"`
	Device             string `name:"Device"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
}

type RunInstancesResponse struct {
	RequestId      goaliyun.String
	InstanceIdSets RunInstancesInstanceIdSetList
}

type RunInstancesTagList []RunInstancesTag

func (list *RunInstancesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesTag)
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

type RunInstancesNetworkInterfaceList []RunInstancesNetworkInterface

func (list *RunInstancesNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesNetworkInterface)
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

type RunInstancesDataDiskList []RunInstancesDataDisk

func (list *RunInstancesDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesDataDisk)
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

type RunInstancesInstanceIdSetList []goaliyun.String

func (list *RunInstancesInstanceIdSetList) UnmarshalJSON(data []byte) error {
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
