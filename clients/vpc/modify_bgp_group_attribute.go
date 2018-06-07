package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyBgpGroupAttributeRequest struct {
	AuthKey              string `position:"Query" name:"AuthKey"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerAsn              int64  `position:"Query" name:"PeerAsn"`
	IsFakeAsn            string `position:"Query" name:"IsFakeAsn"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyBgpGroupAttributeRequest) Invoke(client goaliyun.Client) (*ModifyBgpGroupAttributeResponse, error) {
	resp := &ModifyBgpGroupAttributeResponse{}
	err := client.Request("vpc", "ModifyBgpGroupAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyBgpGroupAttributeResponse struct {
	RequestId goaliyun.String
}
