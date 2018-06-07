package cdn

import (
	"github.com/morlay/goaliyun"
)

type SetUserBlackListRequest struct {
	ConfigUrl     string `position:"Query" name:"ConfigUrl"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *SetUserBlackListRequest) Invoke(client goaliyun.Client) (*SetUserBlackListResponse, error) {
	resp := &SetUserBlackListResponse{}
	err := client.Request("cdn", "SetUserBlackList", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetUserBlackListResponse struct {
	RequestId goaliyun.String
}
