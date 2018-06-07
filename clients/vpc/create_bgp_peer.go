package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateBgpPeerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerIpAddress        string `position:"Query" name:"PeerIpAddress"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateBgpPeerRequest) Invoke(client goaliyun.Client) (*CreateBgpPeerResponse, error) {
	resp := &CreateBgpPeerResponse{}
	err := client.Request("vpc", "CreateBgpPeer", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBgpPeerResponse struct {
	RequestId goaliyun.String
	BgpPeerId goaliyun.String
}
