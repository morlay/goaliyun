package dcdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteDcdnDomainRequest struct {
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	OwnerAccount    string `position:"Query" name:"OwnerAccount"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteDcdnDomainRequest) Invoke(client goaliyun.Client) (*DeleteDcdnDomainResponse, error) {
	resp := &DeleteDcdnDomainResponse{}
	err := client.Request("dcdn", "DeleteDcdnDomain", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDcdnDomainResponse struct {
	RequestId goaliyun.String
}
