package httpdns

import (
	"github.com/morlay/goaliyun"
)

type DeleteDomainRequest struct {
	AccountId  string `position:"Query" name:"AccountId"`
	DomainName string `position:"Query" name:"DomainName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteDomainRequest) Invoke(client goaliyun.Client) (*DeleteDomainResponse, error) {
	resp := &DeleteDomainResponse{}
	err := client.Request("httpdns", "DeleteDomain", "2016-02-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDomainResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
}
