package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveRegistrantProfileRequest struct {
	Country                  string `position:"Query" name:"Country"`
	Address                  string `position:"Query" name:"Address"`
	TelArea                  string `position:"Query" name:"TelArea"`
	City                     string `position:"Query" name:"City"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	Telephone                string `position:"Query" name:"Telephone"`
	DefaultRegistrantProfile string `position:"Query" name:"DefaultRegistrantProfile"`
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	TelExt                   string `position:"Query" name:"TelExt"`
	Province                 string `position:"Query" name:"Province"`
	PostalCode               string `position:"Query" name:"PostalCode"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	Lang                     string `position:"Query" name:"Lang"`
	Email                    string `position:"Query" name:"Email"`
	RegistrantName           string `position:"Query" name:"RegistrantName"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *SaveRegistrantProfileRequest) Invoke(client goaliyun.Client) (*SaveRegistrantProfileResponse, error) {
	resp := &SaveRegistrantProfileResponse{}
	err := client.Request("domain-intl", "SaveRegistrantProfile", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveRegistrantProfileResponse struct {
	RequestId           goaliyun.String
	RegistrantProfileId goaliyun.Integer
}
