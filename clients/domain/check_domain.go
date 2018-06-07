package domain

import (
	"github.com/morlay/goaliyun"
)

type CheckDomainRequest struct {
	FeeCurrency string `position:"Query" name:"FeeCurrency"`
	FeePeriod   int64  `position:"Query" name:"FeePeriod"`
	DomainName  string `position:"Query" name:"DomainName"`
	FeeCommand  string `position:"Query" name:"FeeCommand"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CheckDomainRequest) Invoke(client goaliyun.Client) (*CheckDomainResponse, error) {
	resp := &CheckDomainResponse{}
	err := client.Request("domain", "CheckDomain", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckDomainResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
	Avail      goaliyun.String
	Premium    goaliyun.String
	Reason     goaliyun.String
	Price      goaliyun.Integer
}
