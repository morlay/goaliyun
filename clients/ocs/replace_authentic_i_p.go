package ocs

import (
	"github.com/morlay/goaliyun"
)

type ReplaceAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	OldAuthenticIP       string `position:"Query" name:"OldAuthenticIP"`
	NewAuthenticIP       string `position:"Query" name:"NewAuthenticIP"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReplaceAuthenticIPRequest) Invoke(client goaliyun.Client) (*ReplaceAuthenticIPResponse, error) {
	resp := &ReplaceAuthenticIPResponse{}
	err := client.Request("ocs", "ReplaceAuthenticIP", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReplaceAuthenticIPResponse struct {
	RequestId goaliyun.String
}
