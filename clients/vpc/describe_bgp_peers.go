package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBgpPeersRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	BgpPeerId            string `position:"Query" name:"BgpPeerId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBgpPeersRequest) Invoke(client goaliyun.Client) (*DescribeBgpPeersResponse, error) {
	resp := &DescribeBgpPeersResponse{}
	err := client.Request("vpc", "DescribeBgpPeers", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBgpPeersResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	BgpPeers   DescribeBgpPeersBgpPeerList
}

type DescribeBgpPeersBgpPeer struct {
	Name          goaliyun.String
	Description   goaliyun.String
	BgpPeerId     goaliyun.String
	BgpGroupId    goaliyun.String
	PeerIpAddress goaliyun.String
	PeerAsn       goaliyun.String
	AuthKey       goaliyun.String
	RouterId      goaliyun.String
	BgpStatus     goaliyun.String
	Status        goaliyun.String
	Keepalive     goaliyun.String
	LocalAsn      goaliyun.String
	Hold          goaliyun.String
	IsFake        goaliyun.String
	RouteLimit    goaliyun.String
	RegionId      goaliyun.String
}

type DescribeBgpPeersBgpPeerList []DescribeBgpPeersBgpPeer

func (list *DescribeBgpPeersBgpPeerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBgpPeersBgpPeer)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
