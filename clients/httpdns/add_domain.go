package httpdns

import (
	"github.com/morlay/goaliyun"
)

type AddDomainRequest struct {
	AccountId  string `position:"Query" name:"AccountId"`
	DomainName string `position:"Query" name:"DomainName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *AddDomainRequest) Invoke(client goaliyun.Client) (*AddDomainResponse, error) {
	resp := &AddDomainResponse{}
	err := client.Request("httpdns", "AddDomain", "2016-02-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddDomainResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
}
