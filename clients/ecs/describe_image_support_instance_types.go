package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeImageSupportInstanceTypesRequest struct {
	ActionType           string                                       `position:"Query" name:"ActionType"`
	Filters              *DescribeImageSupportInstanceTypesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                        `position:"Query" name:"ResourceOwnerId"`
	ImageId              string                                       `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string                                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64                                        `position:"Query" name:"OwnerId"`
	RegionId             string                                       `position:"Query" name:"RegionId"`
}

func (req *DescribeImageSupportInstanceTypesRequest) Invoke(client goaliyun.Client) (*DescribeImageSupportInstanceTypesResponse, error) {
	resp := &DescribeImageSupportInstanceTypesResponse{}
	err := client.Request("ecs", "DescribeImageSupportInstanceTypes", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeImageSupportInstanceTypesFilter struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type DescribeImageSupportInstanceTypesResponse struct {
	RequestId     goaliyun.String
	RegionId      goaliyun.String
	ImageId       goaliyun.String
	InstanceTypes DescribeImageSupportInstanceTypesInstanceTypeList
}

type DescribeImageSupportInstanceTypesInstanceType struct {
	InstanceTypeId     goaliyun.String
	CpuCoreCount       goaliyun.Integer
	MemorySize         goaliyun.Float
	InstanceTypeFamily goaliyun.String
}

type DescribeImageSupportInstanceTypesFilterList []DescribeImageSupportInstanceTypesFilter

func (list *DescribeImageSupportInstanceTypesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSupportInstanceTypesFilter)
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

type DescribeImageSupportInstanceTypesInstanceTypeList []DescribeImageSupportInstanceTypesInstanceType

func (list *DescribeImageSupportInstanceTypesInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSupportInstanceTypesInstanceType)
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
