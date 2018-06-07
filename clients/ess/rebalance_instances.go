package ess

import (
	"github.com/morlay/goaliyun"
)

type RebalanceInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RebalanceInstancesRequest) Invoke(client goaliyun.Client) (*RebalanceInstancesResponse, error) {
	resp := &RebalanceInstancesResponse{}
	err := client.Request("ess", "RebalanceInstances", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RebalanceInstancesResponse struct {
	ScalingActivityId goaliyun.String
	RequestId         goaliyun.String
}
