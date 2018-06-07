package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceTypeFamiliesRequest struct {
	Generation           string `position:"Query" name:"Generation"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceTypeFamiliesRequest) Invoke(client goaliyun.Client) (*DescribeInstanceTypeFamiliesResponse, error) {
	resp := &DescribeInstanceTypeFamiliesResponse{}
	err := client.Request("ecs", "DescribeInstanceTypeFamilies", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceTypeFamiliesResponse struct {
	RequestId            goaliyun.String
	InstanceTypeFamilies DescribeInstanceTypeFamiliesInstanceTypeFamilyList
}

type DescribeInstanceTypeFamiliesInstanceTypeFamily struct {
	InstanceTypeFamilyId goaliyun.String
	Generation           goaliyun.String
}

type DescribeInstanceTypeFamiliesInstanceTypeFamilyList []DescribeInstanceTypeFamiliesInstanceTypeFamily

func (list *DescribeInstanceTypeFamiliesInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceTypeFamiliesInstanceTypeFamily)
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
