package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForTransferProhibitionLockRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForTransferProhibitionLockRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForTransferProhibitionLockResponse, error) {
	resp := &SaveSingleTaskForTransferProhibitionLockResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForTransferProhibitionLock", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForTransferProhibitionLockResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
