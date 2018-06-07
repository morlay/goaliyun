package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeInstancePhysicalAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancePhysicalAttributeRequest) Invoke(client goaliyun.Client) (*DescribeInstancePhysicalAttributeResponse, error) {
	resp := &DescribeInstancePhysicalAttributeResponse{}
	err := client.Request("ecs", "DescribeInstancePhysicalAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancePhysicalAttributeResponse struct {
	RequestId        goaliyun.String
	InstanceId       goaliyun.String
	VlanId           goaliyun.String
	NodeControllerId goaliyun.String
	RackId           goaliyun.String
}
