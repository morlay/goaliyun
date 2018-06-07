package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest struct {
	Country                  string                                                               `position:"Query" name:"Country"`
	IdentityCredentialType   string                                                               `position:"Query" name:"IdentityCredentialType"`
	Address                  string                                                               `position:"Query" name:"Address"`
	TelArea                  string                                                               `position:"Query" name:"TelArea"`
	City                     string                                                               `position:"Query" name:"City"`
	ZhAddress                string                                                               `position:"Query" name:"ZhAddress"`
	RegistrantType           string                                                               `position:"Query" name:"RegistrantType"`
	DomainNames              *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	IdentityCredential       string                                                               `position:"Body" name:"IdentityCredential"`
	Telephone                string                                                               `position:"Query" name:"Telephone"`
	TransferOutProhibited    string                                                               `position:"Query" name:"TransferOutProhibited"`
	ZhCity                   string                                                               `position:"Query" name:"ZhCity"`
	ZhProvince               string                                                               `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string                                                               `position:"Query" name:"RegistrantOrganization"`
	TelExt                   string                                                               `position:"Query" name:"TelExt"`
	Province                 string                                                               `position:"Query" name:"Province"`
	ZhRegistrantName         string                                                               `position:"Query" name:"ZhRegistrantName"`
	PostalCode               string                                                               `position:"Query" name:"PostalCode"`
	UserClientIp             string                                                               `position:"Query" name:"UserClientIp"`
	Lang                     string                                                               `position:"Query" name:"Lang"`
	IdentityCredentialNo     string                                                               `position:"Query" name:"IdentityCredentialNo"`
	Email                    string                                                               `position:"Query" name:"Email"`
	RegistrantName           string                                                               `position:"Query" name:"RegistrantName"`
	ZhRegistrantOrganization string                                                               `position:"Query" name:"ZhRegistrantOrganization"`
	RegionId                 string                                                               `position:"Query" name:"RegionId"`
}

func (req *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) Invoke(client goaliyun.Client) (*SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, error) {
	resp := &SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse{}
	err := client.Request("domain", "SaveTaskForUpdatingRegistrantInfoByIdentityCredential", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialDomainNameList []string

func (list *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialDomainNameList) UnmarshalJSON(data []byte) error {
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
