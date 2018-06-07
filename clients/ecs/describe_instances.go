package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	InnerIpAddresses     string `position:"Query" name:"InnerIpAddresses"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	PrivateIpAddresses   string `position:"Query" name:"PrivateIpAddresses"`
	HpcClusterId         string `position:"Query" name:"HpcClusterId"`
	Filter2Value         string `position:"Query" name:"Filter.2.Value"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	LockReason           string `position:"Query" name:"LockReason"`
	Filter1Key           string `position:"Query" name:"Filter.1.Key"`
	DeviceAvailable      string `position:"Query" name:"DeviceAvailable"`
	Filter3Value         string `position:"Query" name:"Filter.3.Value"`
	DryRun               string `position:"Query" name:"DryRun"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	Filter1Value         string `position:"Query" name:"Filter.1.Value"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
	Status               string `position:"Query" name:"Status"`
	ImageId              string `position:"Query" name:"ImageId"`
	Filter4Value         string `position:"Query" name:"Filter.4.Value"`
	IoOptimized          string `position:"Query" name:"IoOptimized"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Filter4Key           string `position:"Query" name:"Filter.4.Key"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RdmaIpAddresses      string `position:"Query" name:"RdmaIpAddresses"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	PublicIpAddresses    string `position:"Query" name:"PublicIpAddresses"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceTypeFamily   string `position:"Query" name:"InstanceTypeFamily"`
	Filter2Key           string `position:"Query" name:"Filter.2.Key"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	EipAddresses         string `position:"Query" name:"EipAddresses"`
	VpcId                string `position:"Query" name:"VpcId"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Filter3Key           string `position:"Query" name:"Filter.3.Key"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesRequest) Invoke(client goaliyun.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("ecs", "DescribeInstances", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Instances  DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId              goaliyun.String
	InstanceName            goaliyun.String
	Description             goaliyun.String
	ImageId                 goaliyun.String
	OSName                  goaliyun.String
	OSType                  goaliyun.String
	RegionId                goaliyun.String
	ZoneId                  goaliyun.String
	ClusterId               goaliyun.String
	InstanceType            goaliyun.String
	Cpu                     goaliyun.Integer
	Memory                  goaliyun.Integer
	HostName                goaliyun.String
	Status                  goaliyun.String
	SerialNumber            goaliyun.String
	InternetChargeType      goaliyun.String
	InternetMaxBandwidthIn  goaliyun.Integer
	InternetMaxBandwidthOut goaliyun.Integer
	VlanId                  goaliyun.String
	CreationTime            goaliyun.String
	StartTime               goaliyun.String
	InstanceNetworkType     goaliyun.String
	InstanceChargeType      goaliyun.String
	SaleCycle               goaliyun.String
	ExpiredTime             goaliyun.String
	AutoReleaseTime         goaliyun.String
	IoOptimized             bool
	DeviceAvailable         bool
	InstanceTypeFamily      goaliyun.String
	LocalStorageCapacity    goaliyun.Integer
	LocalStorageAmount      goaliyun.Integer
	GPUAmount               goaliyun.Integer
	GPUSpec                 goaliyun.String
	SpotStrategy            goaliyun.String
	SpotPriceLimit          goaliyun.Float
	ResourceGroupId         goaliyun.String
	KeyPairName             goaliyun.String
	Recyclable              bool
	HpcClusterId            goaliyun.String
	StoppedMode             goaliyun.String
	NetworkInterfaces       DescribeInstancesNetworkInterfaceList
	OperationLocks          DescribeInstancesLockReasonList
	Tags                    DescribeInstancesTagList
	SecurityGroupIds        DescribeInstancesSecurityGroupIdList
	PublicIpAddress         DescribeInstancesPublicIpAddresList
	InnerIpAddress          DescribeInstancesInnerIpAddresList
	RdmaIpAddress           DescribeInstancesRdmaIpAddresList
	VpcAttributes           DescribeInstancesVpcAttributes
	EipAddress              DescribeInstancesEipAddress
	DedicatedHostAttribute  DescribeInstancesDedicatedHostAttribute
}

type DescribeInstancesNetworkInterface struct {
	NetworkInterfaceId goaliyun.String
	MacAddress         goaliyun.String
	PrimaryIpAddress   goaliyun.String
}

type DescribeInstancesLockReason struct {
	LockReason goaliyun.String
	LockMsg    goaliyun.String
}

type DescribeInstancesTag struct {
	TagKey   goaliyun.String
	TagValue goaliyun.String
}

type DescribeInstancesVpcAttributes struct {
	VpcId            goaliyun.String
	VSwitchId        goaliyun.String
	NatIpAddress     goaliyun.String
	PrivateIpAddress DescribeInstancesPrivateIpAddresList
}

type DescribeInstancesEipAddress struct {
	AllocationId         goaliyun.String
	IpAddress            goaliyun.String
	Bandwidth            goaliyun.Integer
	InternetChargeType   goaliyun.String
	IsSupportUnassociate bool
}

type DescribeInstancesDedicatedHostAttribute struct {
	DedicatedHostId   goaliyun.String
	DedicatedHostName goaliyun.String
}

type DescribeInstancesInstanceList []DescribeInstancesInstance

func (list *DescribeInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInstance)
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

type DescribeInstancesNetworkInterfaceList []DescribeInstancesNetworkInterface

func (list *DescribeInstancesNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesNetworkInterface)
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

type DescribeInstancesLockReasonList []DescribeInstancesLockReason

func (list *DescribeInstancesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesLockReason)
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

type DescribeInstancesTagList []DescribeInstancesTag

func (list *DescribeInstancesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesTag)
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

type DescribeInstancesSecurityGroupIdList []goaliyun.String

func (list *DescribeInstancesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesPublicIpAddresList []goaliyun.String

func (list *DescribeInstancesPublicIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesInnerIpAddresList []goaliyun.String

func (list *DescribeInstancesInnerIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesRdmaIpAddresList []goaliyun.String

func (list *DescribeInstancesRdmaIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesPrivateIpAddresList []goaliyun.String

func (list *DescribeInstancesPrivateIpAddresList) UnmarshalJSON(data []byte) error {
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
