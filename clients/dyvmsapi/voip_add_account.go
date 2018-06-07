package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type VoipAddAccountRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DeviceId             string `position:"Query" name:"DeviceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *VoipAddAccountRequest) Invoke(client goaliyun.Client) (*VoipAddAccountResponse, error) {
	resp := &VoipAddAccountResponse{}
	err := client.Request("dyvmsapi", "VoipAddAccount", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VoipAddAccountResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
