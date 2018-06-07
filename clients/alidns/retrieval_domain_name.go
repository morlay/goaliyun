package alidns

import (
	"github.com/morlay/goaliyun"
)

type RetrievalDomainNameRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *RetrievalDomainNameRequest) Invoke(client goaliyun.Client) (*RetrievalDomainNameResponse, error) {
	resp := &RetrievalDomainNameResponse{}
	err := client.Request("alidns", "RetrievalDomainName", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RetrievalDomainNameResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
	WhoisEmail goaliyun.String
}
