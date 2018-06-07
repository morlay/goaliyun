package ecs

import (
	"github.com/morlay/goaliyun"
)

type RebootInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ForceStop            string `position:"Query" name:"ForceStop"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RebootInstanceRequest) Invoke(client goaliyun.Client) (*RebootInstanceResponse, error) {
	resp := &RebootInstanceResponse{}
	err := client.Request("ecs", "RebootInstance", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RebootInstanceResponse struct {
	RequestId goaliyun.String
}
