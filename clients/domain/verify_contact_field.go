package domain

import (
	"github.com/morlay/goaliyun"
)

type VerifyContactFieldRequest struct {
	Country                  string `position:"Query" name:"Country"`
	Address                  string `position:"Query" name:"Address"`
	TelArea                  string `position:"Query" name:"TelArea"`
	City                     string `position:"Query" name:"City"`
	ZhAddress                string `position:"Query" name:"ZhAddress"`
	RegistrantType           string `position:"Query" name:"RegistrantType"`
	Telephone                string `position:"Query" name:"Telephone"`
	ZhCity                   string `position:"Query" name:"ZhCity"`
	ZhProvince               string `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	TelExt                   string `position:"Query" name:"TelExt"`
	Province                 string `position:"Query" name:"Province"`
	ZhRegistrantName         string `position:"Query" name:"ZhRegistrantName"`
	PostalCode               string `position:"Query" name:"PostalCode"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	Lang                     string `position:"Query" name:"Lang"`
	Email                    string `position:"Query" name:"Email"`
	RegistrantName           string `position:"Query" name:"RegistrantName"`
	ZhRegistrantOrganization string `position:"Query" name:"ZhRegistrantOrganization"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *VerifyContactFieldRequest) Invoke(client goaliyun.Client) (*VerifyContactFieldResponse, error) {
	resp := &VerifyContactFieldResponse{}
	err := client.Request("domain", "VerifyContactField", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VerifyContactFieldResponse struct {
	RequestId goaliyun.String
}
