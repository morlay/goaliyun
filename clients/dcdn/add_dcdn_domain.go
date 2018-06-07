package dcdn

import (
	"github.com/morlay/goaliyun"
)

type AddDcdnDomainRequest struct {
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	Sources         string `position:"Query" name:"Sources"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	OwnerAccount    string `position:"Query" name:"OwnerAccount"`
	Scope           string `position:"Query" name:"Scope"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	CheckUrl        string `position:"Query" name:"CheckUrl"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AddDcdnDomainRequest) Invoke(client goaliyun.Client) (*AddDcdnDomainResponse, error) {
	resp := &AddDcdnDomainResponse{}
	err := client.Request("dcdn", "AddDcdnDomain", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddDcdnDomainResponse struct {
	RequestId goaliyun.String
}
