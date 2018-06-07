package cdn

import (
	"github.com/morlay/goaliyun"
)

type ClearUserBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ClearUserBlackListRequest) Invoke(client goaliyun.Client) (*ClearUserBlackListResponse, error) {
	resp := &ClearUserBlackListResponse{}
	err := client.Request("cdn", "ClearUserBlackList", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ClearUserBlackListResponse struct {
	RequestId goaliyun.String
}
