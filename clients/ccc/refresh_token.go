package ccc

import (
	"github.com/morlay/goaliyun"
)

type RefreshTokenRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RefreshTokenRequest) Invoke(client goaliyun.Client) (*RefreshTokenResponse, error) {
	resp := &RefreshTokenResponse{}
	err := client.Request("ccc", "RefreshToken", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefreshTokenResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Token          RefreshTokenToken
}

type RefreshTokenToken struct {
	Signature goaliyun.String
	SignData  goaliyun.String
}
