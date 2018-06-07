package cf

import (
	"github.com/morlay/goaliyun"
)

type AuthenticateRequest struct {
	Sig       string `position:"Query" name:"Sig"`
	RemoteIp  string `position:"Query" name:"RemoteIp"`
	AccessKey string `position:"Query" name:"AccessKey"`
	AppKey    string `position:"Query" name:"AppKey"`
	SessionId string `position:"Query" name:"SessionId"`
	Token     string `position:"Query" name:"Token"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *AuthenticateRequest) Invoke(client goaliyun.Client) (*AuthenticateResponse, error) {
	resp := &AuthenticateResponse{}
	err := client.Request("cf", "Authenticate", "2015-12-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AuthenticateResponse struct {
	RequestId             goaliyun.String
	Success               bool
	Msg                   goaliyun.String
	Code                  goaliyun.Integer
	SigAuthenticateResult AuthenticateSigAuthenticateResult
}

type AuthenticateSigAuthenticateResult struct {
	Timestamp goaliyun.Integer
	Code      goaliyun.Integer
	Msg       goaliyun.String
	Risklevel goaliyun.String
}
