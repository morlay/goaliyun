package ess

import (
	"github.com/morlay/goaliyun"
)

type VerifyAuthenticationRequest struct {
	Uid                  int64  `position:"Query" name:"Uid"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *VerifyAuthenticationRequest) Invoke(client goaliyun.Client) (*VerifyAuthenticationResponse, error) {
	resp := &VerifyAuthenticationResponse{}
	err := client.Request("ess", "VerifyAuthentication", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VerifyAuthenticationResponse struct {
	RequestId goaliyun.String
}
