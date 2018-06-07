package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeUserBusinessBehaviorRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	StatusKey            string `position:"Query" name:"StatusKey"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserBusinessBehaviorRequest) Invoke(client goaliyun.Client) (*DescribeUserBusinessBehaviorResponse, error) {
	resp := &DescribeUserBusinessBehaviorResponse{}
	err := client.Request("ecs", "DescribeUserBusinessBehavior", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserBusinessBehaviorResponse struct {
	RequestId   goaliyun.String
	StatusValue goaliyun.String
}
