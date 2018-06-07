package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetUserDomainBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetUserDomainBlackListRequest) Invoke(client goaliyun.Client) (*SetUserDomainBlackListResponse, error) {
	resp := &SetUserDomainBlackListResponse{}
	err := client.Request("cdn", "SetUserDomainBlackList", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetUserDomainBlackListResponse struct {
	RequestId goaliyun.String
}
