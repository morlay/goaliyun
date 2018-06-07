package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeInstanceVncPasswdRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceVncPasswdRequest) Invoke(client goaliyun.Client) (*DescribeInstanceVncPasswdResponse, error) {
	resp := &DescribeInstanceVncPasswdResponse{}
	err := client.Request("ecs", "DescribeInstanceVncPasswd", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceVncPasswdResponse struct {
	RequestId goaliyun.String
	VncPasswd goaliyun.String
}
