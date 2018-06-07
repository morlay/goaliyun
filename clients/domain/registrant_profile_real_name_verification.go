package domain

import (
	"github.com/morlay/goaliyun"
)

type RegistrantProfileRealNameVerificationRequest struct {
	IdentityCredentialType string `position:"Query" name:"IdentityCredentialType"`
	UserClientIp           string `position:"Query" name:"UserClientIp"`
	RegistrantProfileID    int64  `position:"Query" name:"RegistrantProfileID"`
	IdentityCredential     string `position:"Body" name:"IdentityCredential"`
	Lang                   string `position:"Query" name:"Lang"`
	IdentityCredentialNo   string `position:"Query" name:"IdentityCredentialNo"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *RegistrantProfileRealNameVerificationRequest) Invoke(client goaliyun.Client) (*RegistrantProfileRealNameVerificationResponse, error) {
	resp := &RegistrantProfileRealNameVerificationResponse{}
	err := client.Request("domain", "RegistrantProfileRealNameVerification", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RegistrantProfileRealNameVerificationResponse struct {
	RequestId goaliyun.String
}
