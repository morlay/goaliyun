package dcdn

import (
	"github.com/morlay/goaliyun"
)

type StopDcdnDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *StopDcdnDomainRequest) Invoke(client goaliyun.Client) (*StopDcdnDomainResponse, error) {
	resp := &StopDcdnDomainResponse{}
	err := client.Request("dcdn", "StopDcdnDomain", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopDcdnDomainResponse struct {
	RequestId goaliyun.String
}
