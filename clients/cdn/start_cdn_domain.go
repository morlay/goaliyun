package cdn

import (
	"github.com/morlay/goaliyun"
)

type StartCdnDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *StartCdnDomainRequest) Invoke(client goaliyun.Client) (*StartCdnDomainResponse, error) {
	resp := &StartCdnDomainResponse{}
	err := client.Request("cdn", "StartCdnDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartCdnDomainResponse struct {
	RequestId goaliyun.String
}
