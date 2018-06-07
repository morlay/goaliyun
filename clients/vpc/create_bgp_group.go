package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateBgpGroupRequest struct {
	AuthKey              string `position:"Query" name:"AuthKey"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerAsn              int64  `position:"Query" name:"PeerAsn"`
	IsFakeAsn            string `position:"Query" name:"IsFakeAsn"`
	RouterId             string `position:"Query" name:"RouterId"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateBgpGroupRequest) Invoke(client goaliyun.Client) (*CreateBgpGroupResponse, error) {
	resp := &CreateBgpGroupResponse{}
	err := client.Request("vpc", "CreateBgpGroup", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBgpGroupResponse struct {
	RequestId  goaliyun.String
	BgpGroupId goaliyun.String
}
