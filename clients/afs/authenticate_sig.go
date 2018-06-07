package afs

import (
	"github.com/morlay/goaliyun"
)

type AuthenticateSigRequest struct {
	Sig             string `position:"Query" name:"Sig"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RemoteIp        string `position:"Query" name:"RemoteIp"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	AppKey          string `position:"Query" name:"AppKey"`
	SessionId       string `position:"Query" name:"SessionId"`
	Token           string `position:"Query" name:"Token"`
	Scene           string `position:"Query" name:"Scene"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AuthenticateSigRequest) Invoke(client goaliyun.Client) (*AuthenticateSigResponse, error) {
	resp := &AuthenticateSigResponse{}
	err := client.Request("afs", "AuthenticateSig", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AuthenticateSigResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.Integer
	Msg       goaliyun.String
	RiskLevel goaliyun.String
	Detail    goaliyun.String
}
