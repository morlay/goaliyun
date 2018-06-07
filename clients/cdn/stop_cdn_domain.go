package cdn

import (
	"github.com/morlay/goaliyun"
)

type StopCdnDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *StopCdnDomainRequest) Invoke(client goaliyun.Client) (*StopCdnDomainResponse, error) {
	resp := &StopCdnDomainResponse{}
	err := client.Request("cdn", "StopCdnDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopCdnDomainResponse struct {
	RequestId goaliyun.String
}
