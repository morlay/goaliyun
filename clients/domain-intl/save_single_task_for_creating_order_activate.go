package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCreatingOrderActivateRequest struct {
	Country                 string `position:"Query" name:"Country"`
	SubscriptionDuration    int64  `position:"Query" name:"SubscriptionDuration"`
	Address                 string `position:"Query" name:"Address"`
	PermitPremiumActivation string `position:"Query" name:"PermitPremiumActivation"`
	TelArea                 string `position:"Query" name:"TelArea"`
	City                    string `position:"Query" name:"City"`
	Dns2                    string `position:"Query" name:"Dns.2"`
	Dns1                    string `position:"Query" name:"Dns.1"`
	DomainName              string `position:"Query" name:"DomainName"`
	RegistrantProfileId     int64  `position:"Query" name:"RegistrantProfileId"`
	Telephone               string `position:"Query" name:"Telephone"`
	AliyunDns               string `position:"Query" name:"AliyunDns"`
	RegistrantOrganization  string `position:"Query" name:"RegistrantOrganization"`
	TelExt                  string `position:"Query" name:"TelExt"`
	Province                string `position:"Query" name:"Province"`
	PostalCode              string `position:"Query" name:"PostalCode"`
	UserClientIp            string `position:"Query" name:"UserClientIp"`
	EnableDomainProxy       string `position:"Query" name:"EnableDomainProxy"`
	Lang                    string `position:"Query" name:"Lang"`
	Email                   string `position:"Query" name:"Email"`
	RegistrantName          string `position:"Query" name:"RegistrantName"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCreatingOrderActivateRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCreatingOrderActivateResponse, error) {
	resp := &SaveSingleTaskForCreatingOrderActivateResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForCreatingOrderActivate", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCreatingOrderActivateResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
