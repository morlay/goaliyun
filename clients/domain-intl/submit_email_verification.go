package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitEmailVerificationRequest struct {
	SendIfExist  string `position:"Query" name:"SendIfExist"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SubmitEmailVerificationRequest) Invoke(client goaliyun.Client) (*SubmitEmailVerificationResponse, error) {
	resp := &SubmitEmailVerificationResponse{}
	err := client.Request("domain-intl", "SubmitEmailVerification", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitEmailVerificationResponse struct {
	RequestId   goaliyun.String
	SuccessList SubmitEmailVerificationSendResultList
	FailList    SubmitEmailVerificationSendResultList
	ExistList   SubmitEmailVerificationSendResultList
}

type SubmitEmailVerificationSendResult struct {
	Email   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type SubmitEmailVerificationSendResultList []SubmitEmailVerificationSendResult

func (list *SubmitEmailVerificationSendResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEmailVerificationSendResult)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
