package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type QueryContactInfoRequest struct {
	ContactType  string `position:"Query" name:"ContactType"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *QueryContactInfoRequest) Invoke(client goaliyun.Client) (*QueryContactInfoResponse, error) {
	resp := &QueryContactInfoResponse{}
	err := client.Request("domain-intl", "QueryContactInfo", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryContactInfoResponse struct {
	RequestId              goaliyun.String
	CreateDate             goaliyun.String
	RegistrantName         goaliyun.String
	RegistrantOrganization goaliyun.String
	Country                goaliyun.String
	Province               goaliyun.String
	City                   goaliyun.String
	Address                goaliyun.String
	Email                  goaliyun.String
	PostalCode             goaliyun.String
	TelArea                goaliyun.String
	Telephone              goaliyun.String
	TelExt                 goaliyun.String
}
