package dcdn

import (
	"github.com/morlay/goaliyun"
)

type UpdateDcdnDomainRequest struct {
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	Sources         string `position:"Query" name:"Sources"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *UpdateDcdnDomainRequest) Invoke(client goaliyun.Client) (*UpdateDcdnDomainResponse, error) {
	resp := &UpdateDcdnDomainResponse{}
	err := client.Request("dcdn", "UpdateDcdnDomain", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateDcdnDomainResponse struct {
	RequestId goaliyun.String
}
