package ocs

import (
	"github.com/morlay/goaliyun"
)

type ActivateInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ActivateInstanceRequest) Invoke(client goaliyun.Client) (*ActivateInstanceResponse, error) {
	resp := &ActivateInstanceResponse{}
	err := client.Request("ocs", "ActivateInstance", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActivateInstanceResponse struct {
	RequestId goaliyun.String
}
