package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SaveBatchTaskForCreatingOrderActivateRequest struct {
	OrderActivateParams *SaveBatchTaskForCreatingOrderActivateOrderActivateParamList `position:"Query" type:"Repeated" name:"OrderActivateParam"`
	UserClientIp        string                                                       `position:"Query" name:"UserClientIp"`
	Lang                string                                                       `position:"Query" name:"Lang"`
	RegionId            string                                                       `position:"Query" name:"RegionId"`
}

func (req *SaveBatchTaskForCreatingOrderActivateRequest) Invoke(client goaliyun.Client) (*SaveBatchTaskForCreatingOrderActivateResponse, error) {
	resp := &SaveBatchTaskForCreatingOrderActivateResponse{}
	err := client.Request("domain", "SaveBatchTaskForCreatingOrderActivate", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveBatchTaskForCreatingOrderActivateOrderActivateParam struct {
	DomainName               string `name:"DomainName"`
	SubscriptionDuration     int64  `name:"SubscriptionDuration"`
	RegistrantProfileId      int64  `name:"RegistrantProfileId"`
	EnableDomainProxy        string `name:"EnableDomainProxy"`
	PermitPremiumActivation  string `name:"PermitPremiumActivation"`
	AliyunDns                string `name:"AliyunDns"`
	Dns1                     string `name:"Dns.1"`
	Dns2                     string `name:"Dns.2"`
	ZhCity                   string `name:"ZhCity"`
	ZhRegistrantOrganization string `name:"ZhRegistrantOrganization"`
	Country                  string `name:"Country"`
	ZhRegistrantName         string `name:"ZhRegistrantName"`
	ZhProvince               string `name:"ZhProvince"`
	ZhAddress                string `name:"ZhAddress"`
	City                     string `name:"City"`
	RegistrantOrganization   string `name:"RegistrantOrganization"`
	RegistrantName           string `name:"RegistrantName"`
	Province                 string `name:"Province"`
	Address                  string `name:"Address"`
	Email                    string `name:"Email"`
	PostalCode               string `name:"PostalCode"`
	TelArea                  string `name:"TelArea"`
	Telephone                string `name:"Telephone"`
	TelExt                   string `name:"TelExt"`
	RegistrantType           string `name:"RegistrantType"`
}

type SaveBatchTaskForCreatingOrderActivateResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}

type SaveBatchTaskForCreatingOrderActivateOrderActivateParamList []SaveBatchTaskForCreatingOrderActivateOrderActivateParam

func (list *SaveBatchTaskForCreatingOrderActivateOrderActivateParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SaveBatchTaskForCreatingOrderActivateOrderActivateParam)
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
