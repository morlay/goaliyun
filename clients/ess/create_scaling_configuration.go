package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateScalingConfigurationRequest struct {
	DataDisk3Size               int64                                         `position:"Query" name:"DataDisk.3.Size"`
	ImageId                     string                                        `position:"Query" name:"ImageId"`
	DataDisk1SnapshotId         string                                        `position:"Query" name:"DataDisk.1.SnapshotId"`
	DataDisk3Category           string                                        `position:"Query" name:"DataDisk.3.Category"`
	DataDisk1Device             string                                        `position:"Query" name:"DataDisk.1.Device"`
	ScalingGroupId              string                                        `position:"Query" name:"ScalingGroupId"`
	DataDisk2Device             string                                        `position:"Query" name:"DataDisk.2.Device"`
	InstanceTypes               *CreateScalingConfigurationInstanceTypeList   `position:"Query" type:"Repeated" name:"InstanceType"`
	IoOptimized                 string                                        `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                                        `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     int64                                         `position:"Query" name:"InternetMaxBandwidthOut"`
	SecurityEnhancementStrategy string                                        `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                                        `position:"Query" name:"KeyPairName"`
	SpotPriceLimits             *CreateScalingConfigurationSpotPriceLimitList `position:"Query" type:"Repeated" name:"SpotPriceLimit"`
	SystemDiskCategory          string                                        `position:"Query" name:"SystemDiskCategory"`
	UserData                    string                                        `position:"Query" name:"UserData"`
	DataDisk4Category           string                                        `position:"Query" name:"DataDisk.4.Category"`
	DataDisk2SnapshotId         string                                        `position:"Query" name:"DataDisk.2.SnapshotId"`
	DataDisk4Size               int64                                         `position:"Query" name:"DataDisk.4.Size"`
	InstanceType                string                                        `position:"Query" name:"InstanceType"`
	DataDisk2Category           string                                        `position:"Query" name:"DataDisk.2.Category"`
	DataDisk1Size               int64                                         `position:"Query" name:"DataDisk.1.Size"`
	DataDisk3SnapshotId         string                                        `position:"Query" name:"DataDisk.3.SnapshotId"`
	ResourceOwnerAccount        string                                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                                        `position:"Query" name:"OwnerAccount"`
	DataDisk2Size               int64                                         `position:"Query" name:"DataDisk.2.Size"`
	RamRoleName                 string                                        `position:"Query" name:"RamRoleName"`
	OwnerId                     int64                                         `position:"Query" name:"OwnerId"`
	ScalingConfigurationName    string                                        `position:"Query" name:"ScalingConfigurationName"`
	Tags                        string                                        `position:"Query" name:"Tags"`
	DataDisk2DeleteWithInstance string                                        `position:"Query" name:"DataDisk.2.DeleteWithInstance"`
	SpotStrategy                string                                        `position:"Query" name:"SpotStrategy"`
	DataDisk1Category           string                                        `position:"Query" name:"DataDisk.1.Category"`
	DataDisk3DeleteWithInstance string                                        `position:"Query" name:"DataDisk.3.DeleteWithInstance"`
	LoadBalancerWeight          int64                                         `position:"Query" name:"LoadBalancerWeight"`
	InstanceName                string                                        `position:"Query" name:"InstanceName"`
	SystemDiskSize              int64                                         `position:"Query" name:"SystemDiskSize"`
	DataDisk4SnapshotId         string                                        `position:"Query" name:"DataDisk.4.SnapshotId"`
	DataDisk4Device             string                                        `position:"Query" name:"DataDisk.4.Device"`
	InternetChargeType          string                                        `position:"Query" name:"InternetChargeType"`
	DataDisk3Device             string                                        `position:"Query" name:"DataDisk.3.Device"`
	DataDisk4DeleteWithInstance string                                        `position:"Query" name:"DataDisk.4.DeleteWithInstance"`
	InternetMaxBandwidthIn      int64                                         `position:"Query" name:"InternetMaxBandwidthIn"`
	DataDisk1DeleteWithInstance string                                        `position:"Query" name:"DataDisk.1.DeleteWithInstance"`
	RegionId                    string                                        `position:"Query" name:"RegionId"`
}

func (req *CreateScalingConfigurationRequest) Invoke(client goaliyun.Client) (*CreateScalingConfigurationResponse, error) {
	resp := &CreateScalingConfigurationResponse{}
	err := client.Request("ess", "CreateScalingConfiguration", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateScalingConfigurationResponse struct {
	ScalingConfigurationId goaliyun.String
	RequestId              goaliyun.String
}

type CreateScalingConfigurationInstanceTypeList []string

func (list *CreateScalingConfigurationInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type CreateScalingConfigurationSpotPriceLimitList []string

func (list *CreateScalingConfigurationSpotPriceLimitList) UnmarshalJSON(data []byte) error {
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
