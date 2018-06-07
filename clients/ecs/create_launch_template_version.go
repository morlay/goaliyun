package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateLaunchTemplateVersionRequest struct {
	LaunchTemplateName          string                                           `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId             int64                                            `position:"Query" name:"ResourceOwnerId"`
	SecurityEnhancementStrategy string                                           `position:"Query" name:"SecurityEnhancementStrategy"`
	NetworkType                 string                                           `position:"Query" name:"NetworkType"`
	KeyPairName                 string                                           `position:"Query" name:"KeyPairName"`
	SpotPriceLimit              float64                                          `position:"Query" name:"SpotPriceLimit"`
	ImageOwnerAlias             string                                           `position:"Query" name:"ImageOwnerAlias"`
	ResourceGroupId             string                                           `position:"Query" name:"ResourceGroupId"`
	HostName                    string                                           `position:"Query" name:"HostName"`
	SystemDiskIops              int64                                            `position:"Query" name:"SystemDiskIops"`
	Tags                        *CreateLaunchTemplateVersionTagList              `position:"Query" type:"Repeated" name:"Tag"`
	Period                      int64                                            `position:"Query" name:"Period"`
	LaunchTemplateId            string                                           `position:"Query" name:"LaunchTemplateId"`
	OwnerId                     int64                                            `position:"Query" name:"OwnerId"`
	VSwitchId                   string                                           `position:"Query" name:"VSwitchId"`
	SpotStrategy                string                                           `position:"Query" name:"SpotStrategy"`
	InstanceName                string                                           `position:"Query" name:"InstanceName"`
	InternetChargeType          string                                           `position:"Query" name:"InternetChargeType"`
	ZoneId                      string                                           `position:"Query" name:"ZoneId"`
	InternetMaxBandwidthIn      int64                                            `position:"Query" name:"InternetMaxBandwidthIn"`
	VersionDescription          string                                           `position:"Query" name:"VersionDescription"`
	ImageId                     string                                           `position:"Query" name:"ImageId"`
	IoOptimized                 string                                           `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                                           `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     int64                                            `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                 string                                           `position:"Query" name:"Description"`
	SystemDiskCategory          string                                           `position:"Query" name:"SystemDiskCategory"`
	UserData                    string                                           `position:"Query" name:"UserData"`
	InstanceType                string                                           `position:"Query" name:"InstanceType"`
	InstanceChargeType          string                                           `position:"Query" name:"InstanceChargeType"`
	EnableVmOsConfig            string                                           `position:"Query" name:"EnableVmOsConfig"`
	NetworkInterfaces           *CreateLaunchTemplateVersionNetworkInterfaceList `position:"Query" type:"Repeated" name:"NetworkInterface"`
	ResourceOwnerAccount        string                                           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                                           `position:"Query" name:"OwnerAccount"`
	SystemDiskDiskName          string                                           `position:"Query" name:"SystemDiskDiskName"`
	RamRoleName                 string                                           `position:"Query" name:"RamRoleName"`
	AutoReleaseTime             string                                           `position:"Query" name:"AutoReleaseTime"`
	SpotDuration                int64                                            `position:"Query" name:"SpotDuration"`
	DataDisks                   *CreateLaunchTemplateVersionDataDiskList         `position:"Query" type:"Repeated" name:"DataDisk"`
	SystemDiskSize              int64                                            `position:"Query" name:"SystemDiskSize"`
	VpcId                       string                                           `position:"Query" name:"VpcId"`
	SystemDiskDescription       string                                           `position:"Query" name:"SystemDiskDescription"`
	RegionId                    string                                           `position:"Query" name:"RegionId"`
}

func (req *CreateLaunchTemplateVersionRequest) Invoke(client goaliyun.Client) (*CreateLaunchTemplateVersionResponse, error) {
	resp := &CreateLaunchTemplateVersionResponse{}
	err := client.Request("ecs", "CreateLaunchTemplateVersion", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateLaunchTemplateVersionTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type CreateLaunchTemplateVersionNetworkInterface struct {
	PrimaryIpAddress     string `name:"PrimaryIpAddress"`
	VSwitchId            string `name:"VSwitchId"`
	SecurityGroupId      string `name:"SecurityGroupId"`
	NetworkInterfaceName string `name:"NetworkInterfaceName"`
	Description          string `name:"Description"`
}

type CreateLaunchTemplateVersionDataDisk struct {
	Size               int64  `name:"Size"`
	SnapshotId         string `name:"SnapshotId"`
	Category           string `name:"Category"`
	Encrypted          string `name:"Encrypted"`
	DiskName           string `name:"DiskName"`
	Description        string `name:"Description"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
}

type CreateLaunchTemplateVersionResponse struct {
	RequestId                   goaliyun.String
	LaunchTemplateVersionNumber goaliyun.Integer
}

type CreateLaunchTemplateVersionTagList []CreateLaunchTemplateVersionTag

func (list *CreateLaunchTemplateVersionTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateLaunchTemplateVersionTag)
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

type CreateLaunchTemplateVersionNetworkInterfaceList []CreateLaunchTemplateVersionNetworkInterface

func (list *CreateLaunchTemplateVersionNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateLaunchTemplateVersionNetworkInterface)
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

type CreateLaunchTemplateVersionDataDiskList []CreateLaunchTemplateVersionDataDisk

func (list *CreateLaunchTemplateVersionDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateLaunchTemplateVersionDataDisk)
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
