package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForApprovingTransferOutRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForApprovingTransferOutRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForApprovingTransferOutResponse, error) {
	resp := &SaveSingleTaskForApprovingTransferOutResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForApprovingTransferOut", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForApprovingTransferOutResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
