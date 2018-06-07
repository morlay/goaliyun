package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest struct {
	IdentityCredentialType string                                                                             `position:"Query" name:"IdentityCredentialType"`
	UserClientIp           string                                                                             `position:"Query" name:"UserClientIp"`
	IdentityCredential     string                                                                             `position:"Body" name:"IdentityCredential"`
	DomainNames            *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang                   string                                                                             `position:"Query" name:"Lang"`
	IdentityCredentialNo   string                                                                             `position:"Query" name:"IdentityCredentialNo"`
	RegionId               string                                                                             `position:"Query" name:"RegionId"`
}

func (req *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) Invoke(client goaliyun.Client) (*SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse, error) {
	resp := &SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse{}
	err := client.Request("domain", "SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredential", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialDomainNameList []string

func (list *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialDomainNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
