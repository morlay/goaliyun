package domain

import (
	"github.com/morlay/goaliyun"
)

type TransferInResendMailTokenRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *TransferInResendMailTokenRequest) Invoke(client goaliyun.Client) (*TransferInResendMailTokenResponse, error) {
	resp := &TransferInResendMailTokenResponse{}
	err := client.Request("domain", "TransferInResendMailToken", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TransferInResendMailTokenResponse struct {
	RequestId goaliyun.String
}
