package alidns

import (
	"github.com/morlay/goaliyun"
)

type DeleteDomainRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteDomainRequest) Invoke(client goaliyun.Client) (*DeleteDomainResponse, error) {
	resp := &DeleteDomainResponse{}
	err := client.Request("alidns", "DeleteDomain", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDomainResponse struct {
	RequestId  goaliyun.String
	DomainName goaliyun.String
}
