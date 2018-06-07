package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForDomainNameProxyServiceRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForDomainNameProxyServiceRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForDomainNameProxyServiceResponse, error) {
	resp := &SaveSingleTaskForDomainNameProxyServiceResponse{}
	err := client.Request("domain", "SaveSingleTaskForDomainNameProxyService", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForDomainNameProxyServiceResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
