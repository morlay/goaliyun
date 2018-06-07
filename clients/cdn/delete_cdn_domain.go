package cdn

import (
	"github.com/morlay/goaliyun"
)

type DeleteCdnDomainRequest struct {
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteCdnDomainRequest) Invoke(client goaliyun.Client) (*DeleteCdnDomainResponse, error) {
	resp := &DeleteCdnDomainResponse{}
	err := client.Request("cdn", "DeleteCdnDomain", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCdnDomainResponse struct {
	RequestId goaliyun.String
}
