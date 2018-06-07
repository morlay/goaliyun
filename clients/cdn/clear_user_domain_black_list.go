package cdn

import (
	"github.com/morlay/goaliyun"
)

type ClearUserDomainBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ClearUserDomainBlackListRequest) Invoke(client goaliyun.Client) (*ClearUserDomainBlackListResponse, error) {
	resp := &ClearUserDomainBlackListResponse{}
	err := client.Request("cdn", "ClearUserDomainBlackList", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ClearUserDomainBlackListResponse struct {
	RequestId goaliyun.String
}
