package cdn

import (
	"github.com/morlay/goaliyun"
)

type GetUserDomainBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *GetUserDomainBlackListRequest) Invoke(client goaliyun.Client) (*GetUserDomainBlackListResponse, error) {
	resp := &GetUserDomainBlackListResponse{}
	err := client.Request("cdn", "GetUserDomainBlackList", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserDomainBlackListResponse struct {
	RequestId goaliyun.String
	Bind      goaliyun.String
}
