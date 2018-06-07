package domain

import (
	"github.com/morlay/goaliyun"
)

type QueryRegistrantProfileRealNameVerificationInfoRequest struct {
	FetchImage          string `position:"Query" name:"FetchImage"`
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                string `position:"Query" name:"Lang"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *QueryRegistrantProfileRealNameVerificationInfoRequest) Invoke(client goaliyun.Client) (*QueryRegistrantProfileRealNameVerificationInfoResponse, error) {
	resp := &QueryRegistrantProfileRealNameVerificationInfoResponse{}
	err := client.Request("domain", "QueryRegistrantProfileRealNameVerificationInfo", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryRegistrantProfileRealNameVerificationInfoResponse struct {
	RequestId              goaliyun.String
	SubmissionDate         goaliyun.String
	ModificationDate       goaliyun.String
	IdentityCredential     goaliyun.String
	RegistrantProfileId    goaliyun.Integer
	IdentityCredentialNo   goaliyun.String
	IdentityCredentialType goaliyun.String
}
