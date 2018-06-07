package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryFailReasonForDomainRealNameVerificationRequest struct {
	RealNameVerificationAction string `position:"Query" name:"RealNameVerificationAction"`
	UserClientIp               string `position:"Query" name:"UserClientIp"`
	DomainName                 string `position:"Query" name:"DomainName"`
	Lang                       string `position:"Query" name:"Lang"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *QueryFailReasonForDomainRealNameVerificationRequest) Invoke(client goaliyun.Client) (*QueryFailReasonForDomainRealNameVerificationResponse, error) {
	resp := &QueryFailReasonForDomainRealNameVerificationResponse{}
	err := client.Request("domain", "QueryFailReasonForDomainRealNameVerification", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryFailReasonForDomainRealNameVerificationResponse struct {
	RequestId goaliyun.String
	Data      QueryFailReasonForDomainRealNameVerificationFailRecordList
}

type QueryFailReasonForDomainRealNameVerificationFailRecord struct {
	Date                         goaliyun.String
	FailReason                   goaliyun.String
	DomainNameVerificationStatus goaliyun.String
}

type QueryFailReasonForDomainRealNameVerificationFailRecordList []QueryFailReasonForDomainRealNameVerificationFailRecord

func (list *QueryFailReasonForDomainRealNameVerificationFailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFailReasonForDomainRealNameVerificationFailRecord)
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
