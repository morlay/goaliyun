package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCancelingTransferInRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCancelingTransferInRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCancelingTransferInResponse, error) {
	resp := &SaveSingleTaskForCancelingTransferInResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForCancelingTransferIn", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCancelingTransferInResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
