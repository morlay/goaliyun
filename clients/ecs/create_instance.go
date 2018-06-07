package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateInstanceRequest struct {
	Tag4Value                   string                      `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId             int64                       `position:"Query" name:"ResourceOwnerId"`
	Tag2Key                     string                      `position:"Query" name:"Tag.2.Key"`
	HpcClusterId                string                      `position:"Query" name:"HpcClusterId"`
	Tag3Key                     string                      `position:"Query" name:"Tag.3.Key"`
	SecurityEnhancementStrategy string                      `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                      `position:"Query" name:"KeyPairName"`
	SpotPriceLimit              float64                     `position:"Query" name:"SpotPriceLimit"`
	Tag1Value                   string                      `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId             string                      `position:"Query" name:"ResourceGroupId"`
	HostName                    string                      `position:"Query" name:"HostName"`
	Password                    string                      `position:"Query" name:"Password"`
	AutoRenewPeriod             int64                       `position:"Query" name:"AutoRenewPeriod"`
	NodeControllerId            string                      `position:"Query" name:"NodeControllerId"`
	Period                      int64                       `position:"Query" name:"Period"`
	DryRun                      string                      `position:"Query" name:"DryRun"`
	Tag5Key                     string                      `position:"Query" name:"Tag.5.Key"`
	OwnerId                     int64                       `position:"Query" name:"OwnerId"`
	VSwitchId                   string                      `position:"Query" name:"VSwitchId"`
	PrivateIpAddress            string                      `position:"Query" name:"PrivateIpAddress"`
	SpotStrategy                string                      `position:"Query" name:"SpotStrategy"`
	PeriodUnit                  string                      `position:"Query" name:"PeriodUnit"`
	InstanceName                string                      `position:"Query" name:"InstanceName"`
	AutoRenew                   string                      `position:"Query" name:"AutoRenew"`
	InternetChargeType          string                      `position:"Query" name:"InternetChargeType"`
	ZoneId                      string                      `position:"Query" name:"ZoneId"`
	Tag4Key                     string                      `position:"Query" name:"Tag.4.Key"`
	InternetMaxBandwidthIn      int64                       `position:"Query" name:"InternetMaxBandwidthIn"`
	UseAdditionalService        string                      `position:"Query" name:"UseAdditionalService"`
	ImageId                     string                      `position:"Query" name:"ImageId"`
	ClientToken                 string                      `position:"Query" name:"ClientToken"`
	VlanId                      string                      `position:"Query" name:"VlanId"`
	SpotInterruptionBehavior    string                      `position:"Query" name:"SpotInterruptionBehavior"`
	IoOptimized                 string                      `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                      `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     int64                       `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                 string                      `position:"Query" name:"Description"`
	SystemDiskCategory          string                      `position:"Query" name:"SystemDiskCategory"`
	UserData                    string                      `position:"Query" name:"UserData"`
	InstanceType                string                      `position:"Query" name:"InstanceType"`
	InstanceChargeType          string                      `position:"Query" name:"InstanceChargeType"`
	Tag3Value                   string                      `position:"Query" name:"Tag.3.Value"`
	DeploymentSetId             string                      `position:"Query" name:"DeploymentSetId"`
	InnerIpAddress              string                      `position:"Query" name:"InnerIpAddress"`
	ResourceOwnerAccount        string                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                      `position:"Query" name:"OwnerAccount"`
	SystemDiskDiskName          string                      `position:"Query" name:"SystemDiskDiskName"`
	RamRoleName                 string                      `position:"Query" name:"RamRoleName"`
	DedicatedHostId             string                      `position:"Query" name:"DedicatedHostId"`
	ClusterId                   string                      `position:"Query" name:"ClusterId"`
	DataDisks                   *CreateInstanceDataDiskList `position:"Query" type:"Repeated" name:"DataDisk"`
	Tag5Value                   string                      `position:"Query" name:"Tag.5.Value"`
	Tag1Key                     string                      `position:"Query" name:"Tag.1.Key"`
	SystemDiskSize              int64                       `position:"Query" name:"SystemDiskSize"`
	Tag2Value                   string                      `position:"Query" name:"Tag.2.Value"`
	SystemDiskDescription       string                      `position:"Query" name:"SystemDiskDescription"`
	RegionId                    string                      `position:"Query" name:"RegionId"`
}

func (req *CreateInstanceRequest) Invoke(client goaliyun.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("ecs", "CreateInstance", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateInstanceDataDisk struct {
	Size               int64  `name:"Size"`
	SnapshotId         string `name:"SnapshotId"`
	Category           string `name:"Category"`
	DiskName           string `name:"DiskName"`
	Description        string `name:"Description"`
	Device             string `name:"Device"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
	Encrypted          string `name:"Encrypted"`
}

type CreateInstanceResponse struct {
	RequestId  goaliyun.String
	InstanceId goaliyun.String
}

type CreateInstanceDataDiskList []CreateInstanceDataDisk

func (list *CreateInstanceDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateInstanceDataDisk)
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
