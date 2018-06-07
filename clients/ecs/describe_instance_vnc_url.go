package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeInstanceVncUrlRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceVncUrlRequest) Invoke(client goaliyun.Client) (*DescribeInstanceVncUrlResponse, error) {
	resp := &DescribeInstanceVncUrlResponse{}
	err := client.Request("ecs", "DescribeInstanceVncUrl", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceVncUrlResponse struct {
	RequestId goaliyun.String
	VncUrl    goaliyun.String
}
