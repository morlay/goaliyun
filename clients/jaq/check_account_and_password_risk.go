package jaq

import (
	"github.com/morlay/goaliyun"
)

type CheckAccountAndPasswordRiskRequest struct {
	CallerName   string `position:"Query" name:"CallerName"`
	AccountName  string `position:"Query" name:"AccountName"`
	PasswordHash string `position:"Query" name:"PasswordHash"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CheckAccountAndPasswordRiskRequest) Invoke(client goaliyun.Client) (*CheckAccountAndPasswordRiskResponse, error) {
	resp := &CheckAccountAndPasswordRiskResponse{}
	err := client.Request("jaq", "CheckAccountAndPasswordRisk", "2016-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckAccountAndPasswordRiskResponse struct {
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Data      CheckAccountAndPasswordRiskData
}

type CheckAccountAndPasswordRiskData struct {
	FnalDecision     goaliyun.Integer
	EventId          goaliyun.String
	UserId           goaliyun.String
	FinalScore       goaliyun.Integer
	FinalDesc        goaliyun.String
	Detail           goaliyun.String
	CaptchaCheckData goaliyun.String
}
