package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteBgpPeerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpPeerId            string `position:"Query" name:"BgpPeerId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteBgpPeerRequest) Invoke(client goaliyun.Client) (*DeleteBgpPeerResponse, error) {
	resp := &DeleteBgpPeerResponse{}
	err := client.Request("vpc", "DeleteBgpPeer", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBgpPeerResponse struct {
	RequestId goaliyun.String
}
