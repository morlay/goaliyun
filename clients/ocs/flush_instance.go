package ocs

import (
	"github.com/morlay/goaliyun"
)

type FlushInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *FlushInstanceRequest) Invoke(client goaliyun.Client) (*FlushInstanceResponse, error) {
	resp := &FlushInstanceResponse{}
	err := client.Request("ocs", "FlushInstance", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FlushInstanceResponse struct {
	RequestId goaliyun.String
}
