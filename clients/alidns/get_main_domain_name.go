package alidns

import (
	"github.com/morlay/goaliyun"
)

type GetMainDomainNameRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InputString  string `position:"Query" name:"InputString"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *GetMainDomainNameRequest) Invoke(client goaliyun.Client) (*GetMainDomainNameResponse, error) {
	resp := &GetMainDomainNameResponse{}
	err := client.Request("alidns", "GetMainDomainName", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetMainDomainNameResponse struct {
	RequestId   goaliyun.String
	DomainName  goaliyun.String
	RR          goaliyun.String
	DomainLevel goaliyun.Integer
}
