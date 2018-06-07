package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCreatingOrderTransferRequest struct {
	PermitPremiumTransfer string `position:"Query" name:"PermitPremiumTransfer"`
	AuthorizationCode     string `position:"Query" name:"AuthorizationCode"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	DomainName            string `position:"Query" name:"DomainName"`
	RegistrantProfileId   int64  `position:"Query" name:"RegistrantProfileId"`
	Lang                  string `position:"Query" name:"Lang"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCreatingOrderTransferRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCreatingOrderTransferResponse, error) {
	resp := &SaveSingleTaskForCreatingOrderTransferResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForCreatingOrderTransfer", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCreatingOrderTransferResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
