package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCreatingOrderActivateRequest struct {
	Country                  string `position:"Query" name:"Country"`
	SubscriptionDuration     int64  `position:"Query" name:"SubscriptionDuration"`
	PermitPremiumActivation  string `position:"Query" name:"PermitPremiumActivation"`
	City                     string `position:"Query" name:"City"`
	Dns2                     string `position:"Query" name:"Dns.2"`
	Dns1                     string `position:"Query" name:"Dns.1"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	AliyunDns                string `position:"Query" name:"AliyunDns"`
	ZhCity                   string `position:"Query" name:"ZhCity"`
	TelExt                   string `position:"Query" name:"TelExt"`
	ZhRegistrantName         string `position:"Query" name:"ZhRegistrantName"`
	Province                 string `position:"Query" name:"Province"`
	PostalCode               string `position:"Query" name:"PostalCode"`
	Lang                     string `position:"Query" name:"Lang"`
	Email                    string `position:"Query" name:"Email"`
	ZhRegistrantOrganization string `position:"Query" name:"ZhRegistrantOrganization"`
	Address                  string `position:"Query" name:"Address"`
	TelArea                  string `position:"Query" name:"TelArea"`
	DomainName               string `position:"Query" name:"DomainName"`
	ZhAddress                string `position:"Query" name:"ZhAddress"`
	RegistrantType           string `position:"Query" name:"RegistrantType"`
	Telephone                string `position:"Query" name:"Telephone"`
	ZhProvince               string `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	EnableDomainProxy        string `position:"Query" name:"EnableDomainProxy"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	RegistrantName           string `position:"Query" name:"RegistrantName"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCreatingOrderActivateRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCreatingOrderActivateResponse, error) {
	resp := &SaveSingleTaskForCreatingOrderActivateResponse{}
	err := client.Request("domain", "SaveSingleTaskForCreatingOrderActivate", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCreatingOrderActivateResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
