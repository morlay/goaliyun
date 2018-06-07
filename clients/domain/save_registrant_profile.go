package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveRegistrantProfileRequest struct {
	Country                  string `position:"Query" name:"Country"`
	Address                  string `position:"Query" name:"Address"`
	TelArea                  string `position:"Query" name:"TelArea"`
	City                     string `position:"Query" name:"City"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	ZhAddress                string `position:"Query" name:"ZhAddress"`
	RegistrantType           string `position:"Query" name:"RegistrantType"`
	Telephone                string `position:"Query" name:"Telephone"`
	DefaultRegistrantProfile string `position:"Query" name:"DefaultRegistrantProfile"`
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

func (req *SaveRegistrantProfileRequest) Invoke(client goaliyun.Client) (*SaveRegistrantProfileResponse, error) {
	resp := &SaveRegistrantProfileResponse{}
	err := client.Request("domain", "SaveRegistrantProfile", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveRegistrantProfileResponse struct {
	RequestId           goaliyun.String
	RegistrantProfileId goaliyun.Integer
}
