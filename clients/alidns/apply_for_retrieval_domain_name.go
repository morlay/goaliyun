package alidns

import (
	"github.com/morlay/goaliyun"
)

type ApplyForRetrievalDomainNameRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ApplyForRetrievalDomainNameRequest) Invoke(client goaliyun.Client) (*ApplyForRetrievalDomainNameResponse, error) {
	resp := &ApplyForRetrievalDomainNameResponse{}
	err := client.Request("alidns", "ApplyForRetrievalDomainName", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ApplyForRetrievalDomainNameResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
}
