package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceAttributeRequest) Invoke(client goaliyun.Client) (*DescribeInstanceAttributeResponse, error) {
	resp := &DescribeInstanceAttributeResponse{}
	err := client.Request("ecs", "DescribeInstanceAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceAttributeResponse struct {
	RequestId               goaliyun.String
	InstanceId              goaliyun.String
	InstanceName            goaliyun.String
	ImageId                 goaliyun.String
	RegionId                goaliyun.String
	ZoneId                  goaliyun.String
	ClusterId               goaliyun.String
	InstanceType            goaliyun.String
	Cpu                     goaliyun.Integer
	Memory                  goaliyun.Integer
	HostName                goaliyun.String
	Status                  goaliyun.String
	InternetChargeType      goaliyun.String
	InternetMaxBandwidthIn  goaliyun.Integer
	InternetMaxBandwidthOut goaliyun.Integer
	VlanId                  goaliyun.String
	SerialNumber            goaliyun.String
	CreationTime            goaliyun.String
	Description             goaliyun.String
	InstanceNetworkType     goaliyun.String
	IoOptimized             goaliyun.String
	InstanceChargeType      goaliyun.String
	ExpiredTime             goaliyun.String
	StoppedMode             goaliyun.String
	OperationLocks          DescribeInstanceAttributeLockReasonList
	SecurityGroupIds        DescribeInstanceAttributeSecurityGroupIdList
	PublicIpAddress         DescribeInstanceAttributePublicIpAddresList
	InnerIpAddress          DescribeInstanceAttributeInnerIpAddresList
	VpcAttributes           DescribeInstanceAttributeVpcAttributes
	EipAddress              DescribeInstanceAttributeEipAddress
	DedicatedHostAttribute  DescribeInstanceAttributeDedicatedHostAttribute
}

type DescribeInstanceAttributeLockReason struct {
	LockReason goaliyun.String
}

type DescribeInstanceAttributeVpcAttributes struct {
	VpcId            goaliyun.String
	VSwitchId        goaliyun.String
	NatIpAddress     goaliyun.String
	PrivateIpAddress DescribeInstanceAttributePrivateIpAddresList
}

type DescribeInstanceAttributeEipAddress struct {
	AllocationId       goaliyun.String
	IpAddress          goaliyun.String
	Bandwidth          goaliyun.Integer
	InternetChargeType goaliyun.String
}

type DescribeInstanceAttributeDedicatedHostAttribute struct {
	DedicatedHostId   goaliyun.String
	DedicatedHostName goaliyun.String
}

type DescribeInstanceAttributeLockReasonList []DescribeInstanceAttributeLockReason

func (list *DescribeInstanceAttributeLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAttributeLockReason)
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

type DescribeInstanceAttributeSecurityGroupIdList []goaliyun.String

func (list *DescribeInstanceAttributeSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributePublicIpAddresList []goaliyun.String

func (list *DescribeInstanceAttributePublicIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributeInnerIpAddresList []goaliyun.String

func (list *DescribeInstanceAttributeInnerIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributePrivateIpAddresList []goaliyun.String

func (list *DescribeInstanceAttributePrivateIpAddresList) UnmarshalJSON(data []byte) error {
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
