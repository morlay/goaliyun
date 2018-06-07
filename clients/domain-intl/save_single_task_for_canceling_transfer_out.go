package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCancelingTransferOutRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCancelingTransferOutRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCancelingTransferOutResponse, error) {
	resp := &SaveSingleTaskForCancelingTransferOutResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForCancelingTransferOut", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCancelingTransferOutResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
