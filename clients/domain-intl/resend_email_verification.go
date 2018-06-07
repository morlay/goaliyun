package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ResendEmailVerificationRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ResendEmailVerificationRequest) Invoke(client goaliyun.Client) (*ResendEmailVerificationResponse, error) {
	resp := &ResendEmailVerificationResponse{}
	err := client.Request("domain-intl", "ResendEmailVerification", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResendEmailVerificationResponse struct {
	RequestId   goaliyun.String
	SuccessList ResendEmailVerificationSendResultList
	FailList    ResendEmailVerificationSendResultList
}

type ResendEmailVerificationSendResult struct {
	Email   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type ResendEmailVerificationSendResultList []ResendEmailVerificationSendResult

func (list *ResendEmailVerificationSendResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResendEmailVerificationSendResult)
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
