package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceAutoRenewAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceAutoRenewAttributeRequest) Invoke(client goaliyun.Client) (*DescribeInstanceAutoRenewAttributeResponse, error) {
	resp := &DescribeInstanceAutoRenewAttributeResponse{}
	err := client.Request("ecs", "DescribeInstanceAutoRenewAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceAutoRenewAttributeResponse struct {
	RequestId               goaliyun.String
	InstanceRenewAttributes DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList
}

type DescribeInstanceAutoRenewAttributeInstanceRenewAttribute struct {
	InstanceId       goaliyun.String
	AutoRenewEnabled bool
	Duration         goaliyun.Integer
	PeriodUnit       goaliyun.String
	RenewalStatus    goaliyun.String
}

type DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList []DescribeInstanceAutoRenewAttributeInstanceRenewAttribute

func (list *DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewAttributeInstanceRenewAttribute)
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
