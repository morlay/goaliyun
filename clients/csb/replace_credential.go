package csb

import (
	"github.com/morlay/goaliyun"
)

type ReplaceCredentialRequest struct {
	CredentialId int64  `position:"Query" name:"CredentialId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ReplaceCredentialRequest) Invoke(client goaliyun.Client) (*ReplaceCredentialResponse, error) {
	resp := &ReplaceCredentialResponse{}
	err := client.Request("csb", "ReplaceCredential", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReplaceCredentialResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      ReplaceCredentialData
}

type ReplaceCredentialData struct {
	Credentials ReplaceCredentialCredentials
}

type ReplaceCredentialCredentials struct {
	GmtCreate         goaliyun.Integer
	Id                goaliyun.Integer
	Name              goaliyun.String
	UserId            goaliyun.String
	CurrentCredential ReplaceCredentialCurrentCredential
	NewCredential     ReplaceCredentialNewCredential
}

type ReplaceCredentialCurrentCredential struct {
	AccessKey goaliyun.String
	SecretKey goaliyun.String
}

type ReplaceCredentialNewCredential struct {
	AccessKey goaliyun.String
	SecretKey goaliyun.String
}
