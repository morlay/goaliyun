package csb

import (
	"github.com/morlay/goaliyun"
)

type RenewCredentialsRequest struct {
	CredentialId int64  `position:"Query" name:"CredentialId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *RenewCredentialsRequest) Invoke(client goaliyun.Client) (*RenewCredentialsResponse, error) {
	resp := &RenewCredentialsResponse{}
	err := client.Request("csb", "RenewCredentials", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewCredentialsResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      RenewCredentialsData
}

type RenewCredentialsData struct {
	Credentials RenewCredentialsCredentials
}

type RenewCredentialsCredentials struct {
	GmtCreate         goaliyun.Integer
	Id                goaliyun.Integer
	Name              goaliyun.String
	UserId            goaliyun.String
	CurrentCredential RenewCredentialsCurrentCredential
	NewCredential     RenewCredentialsNewCredential
}

type RenewCredentialsCurrentCredential struct {
	AccessKey goaliyun.String
	SecretKey goaliyun.String
}

type RenewCredentialsNewCredential struct {
	AccessKey goaliyun.String
	SecretKey goaliyun.String
}
