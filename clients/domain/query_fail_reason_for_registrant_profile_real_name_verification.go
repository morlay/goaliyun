package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryFailReasonForRegistrantProfileRealNameVerificationRequest struct {
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	RegistrantProfileID int64  `position:"Query" name:"RegistrantProfileID"`
	Lang                string `position:"Query" name:"Lang"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) Invoke(client goaliyun.Client) (*QueryFailReasonForRegistrantProfileRealNameVerificationResponse, error) {
	resp := &QueryFailReasonForRegistrantProfileRealNameVerificationResponse{}
	err := client.Request("domain", "QueryFailReasonForRegistrantProfileRealNameVerification", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryFailReasonForRegistrantProfileRealNameVerificationResponse struct {
	RequestId goaliyun.String
	Data      QueryFailReasonForRegistrantProfileRealNameVerificationFailRecordList
}

type QueryFailReasonForRegistrantProfileRealNameVerificationFailRecord struct {
	Date       goaliyun.String
	FailReason goaliyun.String
}

type QueryFailReasonForRegistrantProfileRealNameVerificationFailRecordList []QueryFailReasonForRegistrantProfileRealNameVerificationFailRecord

func (list *QueryFailReasonForRegistrantProfileRealNameVerificationFailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFailReasonForRegistrantProfileRealNameVerificationFailRecord)
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
