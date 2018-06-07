package dcdn

import (
	"github.com/morlay/goaliyun"
)

type StartDcdnDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *StartDcdnDomainRequest) Invoke(client goaliyun.Client) (*StartDcdnDomainResponse, error) {
	resp := &StartDcdnDomainResponse{}
	err := client.Request("dcdn", "StartDcdnDomain", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartDcdnDomainResponse struct {
	RequestId goaliyun.String
}
