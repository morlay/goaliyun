package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeScalingConfigurationsRequest struct {
	ScalingConfigurationId6    string `position:"Query" name:"ScalingConfigurationId.6"`
	ScalingConfigurationId7    string `position:"Query" name:"ScalingConfigurationId.7"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingConfigurationId4    string `position:"Query" name:"ScalingConfigurationId.4"`
	ScalingConfigurationId5    string `position:"Query" name:"ScalingConfigurationId.5"`
	ScalingGroupId             string `position:"Query" name:"ScalingGroupId"`
	ScalingConfigurationId8    string `position:"Query" name:"ScalingConfigurationId.8"`
	ScalingConfigurationId9    string `position:"Query" name:"ScalingConfigurationId.9"`
	ScalingConfigurationId10   string `position:"Query" name:"ScalingConfigurationId.10"`
	PageNumber                 int64  `position:"Query" name:"PageNumber"`
	ScalingConfigurationName2  string `position:"Query" name:"ScalingConfigurationName.2"`
	ScalingConfigurationName3  string `position:"Query" name:"ScalingConfigurationName.3"`
	ScalingConfigurationName1  string `position:"Query" name:"ScalingConfigurationName.1"`
	PageSize                   int64  `position:"Query" name:"PageSize"`
	ScalingConfigurationId2    string `position:"Query" name:"ScalingConfigurationId.2"`
	ScalingConfigurationId3    string `position:"Query" name:"ScalingConfigurationId.3"`
	ScalingConfigurationId1    string `position:"Query" name:"ScalingConfigurationId.1"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	ScalingConfigurationName6  string `position:"Query" name:"ScalingConfigurationName.6"`
	ScalingConfigurationName7  string `position:"Query" name:"ScalingConfigurationName.7"`
	ScalingConfigurationName4  string `position:"Query" name:"ScalingConfigurationName.4"`
	ScalingConfigurationName5  string `position:"Query" name:"ScalingConfigurationName.5"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	ScalingConfigurationName8  string `position:"Query" name:"ScalingConfigurationName.8"`
	ScalingConfigurationName9  string `position:"Query" name:"ScalingConfigurationName.9"`
	ScalingConfigurationName10 string `position:"Query" name:"ScalingConfigurationName.10"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *DescribeScalingConfigurationsRequest) Invoke(client goaliyun.Client) (*DescribeScalingConfigurationsResponse, error) {
	resp := &DescribeScalingConfigurationsResponse{}
	err := client.Request("ess", "DescribeScalingConfigurations", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeScalingConfigurationsResponse struct {
	TotalCount            goaliyun.Integer
	PageNumber            goaliyun.Integer
	PageSize              goaliyun.Integer
	RequestId             goaliyun.String
	ScalingConfigurations DescribeScalingConfigurationsScalingConfigurationList
}

type DescribeScalingConfigurationsScalingConfiguration struct {
	ScalingConfigurationId      goaliyun.String
	ScalingConfigurationName    goaliyun.String
	ScalingGroupId              goaliyun.String
	InstanceName                goaliyun.String
	ImageId                     goaliyun.String
	InstanceType                goaliyun.String
	InstanceGeneration          goaliyun.String
	SecurityGroupId             goaliyun.String
	IoOptimized                 goaliyun.String
	InternetChargeType          goaliyun.String
	InternetMaxBandwidthIn      goaliyun.Integer
	InternetMaxBandwidthOut     goaliyun.Integer
	SystemDiskCategory          goaliyun.String
	SystemDiskSize              goaliyun.Integer
	LifecycleState              goaliyun.String
	CreationTime                goaliyun.String
	LoadBalancerWeight          goaliyun.Integer
	UserData                    goaliyun.String
	KeyPairName                 goaliyun.String
	RamRoleName                 goaliyun.String
	DeploymentSetId             goaliyun.String
	SecurityEnhancementStrategy goaliyun.String
	SpotStrategy                goaliyun.String
	DataDisks                   DescribeScalingConfigurationsDataDiskList
	Tags                        DescribeScalingConfigurationsTagList
	SpotPriceLimit              DescribeScalingConfigurationsSpotPriceModelList
	InstanceTypes               DescribeScalingConfigurationsInstanceTypeList
}

type DescribeScalingConfigurationsDataDisk struct {
	Size       goaliyun.Integer
	Category   goaliyun.String
	SnapshotId goaliyun.String
	Device     goaliyun.String
}

type DescribeScalingConfigurationsTag struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type DescribeScalingConfigurationsSpotPriceModel struct {
	InstanceType goaliyun.String
	PriceLimit   goaliyun.Float
}

type DescribeScalingConfigurationsScalingConfigurationList []DescribeScalingConfigurationsScalingConfiguration

func (list *DescribeScalingConfigurationsScalingConfigurationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsScalingConfiguration)
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

type DescribeScalingConfigurationsDataDiskList []DescribeScalingConfigurationsDataDisk

func (list *DescribeScalingConfigurationsDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsDataDisk)
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

type DescribeScalingConfigurationsTagList []DescribeScalingConfigurationsTag

func (list *DescribeScalingConfigurationsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsTag)
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

type DescribeScalingConfigurationsSpotPriceModelList []DescribeScalingConfigurationsSpotPriceModel

func (list *DescribeScalingConfigurationsSpotPriceModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsSpotPriceModel)
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

type DescribeScalingConfigurationsInstanceTypeList []goaliyun.String

func (list *DescribeScalingConfigurationsInstanceTypeList) UnmarshalJSON(data []byte) error {
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
