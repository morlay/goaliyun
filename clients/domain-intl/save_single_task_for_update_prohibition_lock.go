package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForUpdateProhibitionLockRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForUpdateProhibitionLockRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForUpdateProhibitionLockResponse, error) {
	resp := &SaveSingleTaskForUpdateProhibitionLockResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForUpdateProhibitionLock", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForUpdateProhibitionLockResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
