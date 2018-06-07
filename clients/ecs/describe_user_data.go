package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeUserDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserDataRequest) Invoke(client goaliyun.Client) (*DescribeUserDataResponse, error) {
	resp := &DescribeUserDataResponse{}
	err := client.Request("ecs", "DescribeUserData", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserDataResponse struct {
	RequestId  goaliyun.String
	RegionId   goaliyun.String
	InstanceId goaliyun.String
	UserData   goaliyun.String
}
