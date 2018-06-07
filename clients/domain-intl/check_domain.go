package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type CheckDomainRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *CheckDomainRequest) Invoke(client goaliyun.Client) (*CheckDomainResponse, error) {
	resp := &CheckDomainResponse{}
	err := client.Request("domain-intl", "CheckDomain", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckDomainResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
	Avail      goaliyun.String
	Premium    goaliyun.String
}
