package csb

import (
	"github.com/morlay/goaliyun"
)

type CreateCredentialsRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateCredentialsRequest) Invoke(client goaliyun.Client) (*CreateCredentialsResponse, error) {
	resp := &CreateCredentialsResponse{}
	err := client.Request("csb", "CreateCredentials", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateCredentialsResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      CreateCredentialsData
}

type CreateCredentialsData struct {
	Credentials CreateCredentialsCredentials
}

type CreateCredentialsCredentials struct {
	Id                goaliyun.Integer
	CurrentCredential CreateCredentialsCurrentCredential
}

type CreateCredentialsCurrentCredential struct {
	SecretKey goaliyun.String
	AccessKey goaliyun.String
}
