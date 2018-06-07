package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceTypesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceTypeFamily   string `position:"Query" name:"InstanceTypeFamily"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceTypesRequest) Invoke(client goaliyun.Client) (*DescribeInstanceTypesResponse, error) {
	resp := &DescribeInstanceTypesResponse{}
	err := client.Request("ecs", "DescribeInstanceTypes", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceTypesResponse struct {
	RequestId     goaliyun.String
	InstanceTypes DescribeInstanceTypesInstanceTypeList
}

type DescribeInstanceTypesInstanceType struct {
	InstanceTypeId       goaliyun.String
	CpuCoreCount         goaliyun.Integer
	MemorySize           goaliyun.Float
	InstanceTypeFamily   goaliyun.String
	LocalStorageCapacity goaliyun.Integer
	LocalStorageAmount   goaliyun.Integer
	LocalStorageCategory goaliyun.String
	GPUAmount            goaliyun.Integer
	GPUSpec              goaliyun.String
	InitialCredit        goaliyun.Integer
	BaselineCredit       goaliyun.Integer
	EniQuantity          goaliyun.Integer
	InstanceBandwidthRx  goaliyun.Integer
	InstanceBandwidthTx  goaliyun.Integer
	InstancePpsRx        goaliyun.Integer
	InstancePpsTx        goaliyun.Integer
}

type DescribeInstanceTypesInstanceTypeList []DescribeInstanceTypesInstanceType

func (list *DescribeInstanceTypesInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceTypesInstanceType)
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
