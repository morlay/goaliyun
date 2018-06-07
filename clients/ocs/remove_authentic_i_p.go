package ocs

import (
	"github.com/morlay/goaliyun"
)

type RemoveAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	AuthenticIP          string `position:"Query" name:"AuthenticIP"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RemoveAuthenticIPRequest) Invoke(client goaliyun.Client) (*RemoveAuthenticIPResponse, error) {
	resp := &RemoveAuthenticIPResponse{}
	err := client.Request("ocs", "RemoveAuthenticIP", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveAuthenticIPResponse struct {
	RequestId goaliyun.String
}
