package ess

import (
	"github.com/morlay/goaliyun"
)

type VerifyUserRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *VerifyUserRequest) Invoke(client goaliyun.Client) (*VerifyUserResponse, error) {
	resp := &VerifyUserResponse{}
	err := client.Request("ess", "VerifyUser", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VerifyUserResponse struct {
}
