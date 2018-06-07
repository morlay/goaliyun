package cdn

import (
	"github.com/morlay/goaliyun"
)

type ModifyCdnDomainRequest struct {
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	SourcePort      int64  `position:"Query" name:"SourcePort"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	Priorities      string `position:"Query" name:"Priorities"`
	Sources         string `position:"Query" name:"Sources"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	SourceType      string `position:"Query" name:"SourceType"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyCdnDomainRequest) Invoke(client goaliyun.Client) (*ModifyCdnDomainResponse, error) {
	resp := &ModifyCdnDomainResponse{}
	err := client.Request("cdn", "ModifyCdnDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCdnDomainResponse struct {
	RequestId goaliyun.String
}
