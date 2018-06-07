package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type CheckTransferInFeasibilityRequest struct {
	TransferAuthorizationCode string `position:"Query" name:"TransferAuthorizationCode"`
	UserClientIp              string `position:"Query" name:"UserClientIp"`
	DomainName                string `position:"Query" name:"DomainName"`
	Lang                      string `position:"Query" name:"Lang"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *CheckTransferInFeasibilityRequest) Invoke(client goaliyun.Client) (*CheckTransferInFeasibilityResponse, error) {
	resp := &CheckTransferInFeasibilityResponse{}
	err := client.Request("domain-intl", "CheckTransferInFeasibility", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckTransferInFeasibilityResponse struct {
	RequestId   goaliyun.String
	CanTransfer bool
	Code        goaliyun.String
	Message     goaliyun.String
	ProductId   goaliyun.String
}
