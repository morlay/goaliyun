package ocs

import (
	"github.com/morlay/goaliyun"
)

type DeleteInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteInstanceRequest) Invoke(client goaliyun.Client) (*DeleteInstanceResponse, error) {
	resp := &DeleteInstanceResponse{}
	err := client.Request("ocs", "DeleteInstance", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteInstanceResponse struct {
	RequestId goaliyun.String
}
