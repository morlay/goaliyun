package ocs

import (
	"github.com/morlay/goaliyun"
)

type AddAuthenticIPRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	AuthenticIP          string `position:"Query" name:"AuthenticIP"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddAuthenticIPRequest) Invoke(client goaliyun.Client) (*AddAuthenticIPResponse, error) {
	resp := &AddAuthenticIPResponse{}
	err := client.Request("ocs", "AddAuthenticIP", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddAuthenticIPResponse struct {
	RequestId goaliyun.String
}
