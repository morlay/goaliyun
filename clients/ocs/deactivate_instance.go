package ocs

import (
	"github.com/morlay/goaliyun"
)

type DeactivateInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeactivateInstanceRequest) Invoke(client goaliyun.Client) (*DeactivateInstanceResponse, error) {
	resp := &DeactivateInstanceResponse{}
	err := client.Request("ocs", "DeactivateInstance", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeactivateInstanceResponse struct {
	RequestId goaliyun.String
}
