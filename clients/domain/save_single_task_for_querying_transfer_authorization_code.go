package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForQueryingTransferAuthorizationCodeRequest struct {
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, error) {
	resp := &SaveSingleTaskForQueryingTransferAuthorizationCodeResponse{}
	err := client.Request("domain", "SaveSingleTaskForQueryingTransferAuthorizationCode", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
