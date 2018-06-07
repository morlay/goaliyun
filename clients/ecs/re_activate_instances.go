package ecs

import (
	"github.com/morlay/goaliyun"
)

type ReActivateInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReActivateInstancesRequest) Invoke(client goaliyun.Client) (*ReActivateInstancesResponse, error) {
	resp := &ReActivateInstancesResponse{}
	err := client.Request("ecs", "ReActivateInstances", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReActivateInstancesResponse struct {
	RequestId goaliyun.String
}
