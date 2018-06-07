package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type VerifyContactFieldRequest struct {
	Country                string `position:"Query" name:"Country"`
	Address                string `position:"Query" name:"Address"`
	TelArea                string `position:"Query" name:"TelArea"`
	City                   string `position:"Query" name:"City"`
	Telephone              string `position:"Query" name:"Telephone"`
	RegistrantOrganization string `position:"Query" name:"RegistrantOrganization"`
	TelExt                 string `position:"Query" name:"TelExt"`
	Province               string `position:"Query" name:"Province"`
	PostalCode             string `position:"Query" name:"PostalCode"`
	UserClientIp           string `position:"Query" name:"UserClientIp"`
	Lang                   string `position:"Query" name:"Lang"`
	Email                  string `position:"Query" name:"Email"`
	RegistrantName         string `position:"Query" name:"RegistrantName"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *VerifyContactFieldRequest) Invoke(client goaliyun.Client) (*VerifyContactFieldResponse, error) {
	resp := &VerifyContactFieldResponse{}
	err := client.Request("domain-intl", "VerifyContactField", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VerifyContactFieldResponse struct {
	RequestId goaliyun.String
}
