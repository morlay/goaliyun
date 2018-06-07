package domain

import (
	"github.com/morlay/goaliyun"
)

type TransferInRefetchWhoisEmailRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *TransferInRefetchWhoisEmailRequest) Invoke(client goaliyun.Client) (*TransferInRefetchWhoisEmailResponse, error) {
	resp := &TransferInRefetchWhoisEmailResponse{}
	err := client.Request("domain", "TransferInRefetchWhoisEmail", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TransferInRefetchWhoisEmailResponse struct {
	RequestId goaliyun.String
}
