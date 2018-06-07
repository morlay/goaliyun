package alidns

import (
	"github.com/morlay/goaliyun"
)

type ChangeDomainOfDnsProductRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	NewDomain    string `position:"Query" name:"NewDomain"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ChangeDomainOfDnsProductRequest) Invoke(client goaliyun.Client) (*ChangeDomainOfDnsProductResponse, error) {
	resp := &ChangeDomainOfDnsProductResponse{}
	err := client.Request("alidns", "ChangeDomainOfDnsProduct", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ChangeDomainOfDnsProductResponse struct {
	RequestId      goaliyun.String
	OriginalDomain goaliyun.String
}
