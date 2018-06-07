package ocs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceCapacityRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	Capacity             int64  `position:"Query" name:"Capacity"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceCapacityRequest) Invoke(client goaliyun.Client) (*ModifyInstanceCapacityResponse, error) {
	resp := &ModifyInstanceCapacityResponse{}
	err := client.Request("ocs", "ModifyInstanceCapacity", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceCapacityResponse struct {
	RequestId goaliyun.String
}
