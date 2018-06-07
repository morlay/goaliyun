package cdn

import (
	"github.com/morlay/goaliyun"
)

type AddCdnDomainRequest struct {
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	Sources         string `position:"Query" name:"Sources"`
	OwnerAccount    string `position:"Query" name:"OwnerAccount"`
	DomainName      string `position:"Query" name:"DomainName"`
	LiveType        string `position:"Query" name:"LiveType"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourcePort      int64  `position:"Query" name:"SourcePort"`
	Priorities      string `position:"Query" name:"Priorities"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	CdnType         string `position:"Query" name:"CdnType"`
	Scope           string `position:"Query" name:"Scope"`
	SourceType      string `position:"Query" name:"SourceType"`
	CheckUrl        string `position:"Query" name:"CheckUrl"`
	Region          string `position:"Query" name:"Region"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AddCdnDomainRequest) Invoke(client goaliyun.Client) (*AddCdnDomainResponse, error) {
	resp := &AddCdnDomainResponse{}
	err := client.Request("cdn", "AddCdnDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCdnDomainResponse struct {
	RequestId goaliyun.String
}
