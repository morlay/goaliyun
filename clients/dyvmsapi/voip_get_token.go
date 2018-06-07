package dyvmsapi

import (
	"github.com/morlay/goaliyun"
)

type VoipGetTokenRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VoipId               string `position:"Query" name:"VoipId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DeviceId             string `position:"Query" name:"DeviceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *VoipGetTokenRequest) Invoke(client goaliyun.Client) (*VoipGetTokenResponse, error) {
	resp := &VoipGetTokenResponse{}
	err := client.Request("dyvmsapi", "VoipGetToken", "2017-05-25", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VoipGetTokenResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
