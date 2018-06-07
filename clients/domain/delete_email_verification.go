package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteEmailVerificationRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteEmailVerificationRequest) Invoke(client goaliyun.Client) (*DeleteEmailVerificationResponse, error) {
	resp := &DeleteEmailVerificationResponse{}
	err := client.Request("domain", "DeleteEmailVerification", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteEmailVerificationResponse struct {
	RequestId   goaliyun.String
	SuccessList DeleteEmailVerificationSendResultList
	FailList    DeleteEmailVerificationSendResultList
}

type DeleteEmailVerificationSendResult struct {
	Email   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type DeleteEmailVerificationSendResultList []DeleteEmailVerificationSendResult

func (list *DeleteEmailVerificationSendResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteEmailVerificationSendResult)
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
