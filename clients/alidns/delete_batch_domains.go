package alidns

import (
	"github.com/morlay/goaliyun"
)

type DeleteBatchDomainsRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Domains      string `position:"Query" name:"Domains"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteBatchDomainsRequest) Invoke(client goaliyun.Client) (*DeleteBatchDomainsResponse, error) {
	resp := &DeleteBatchDomainsResponse{}
	err := client.Request("alidns", "DeleteBatchDomains", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBatchDomainsResponse struct {
	RequestId goaliyun.String
	TraceId   goaliyun.String
}
