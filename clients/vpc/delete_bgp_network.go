package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteBgpNetworkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DstCidrBlock         string `position:"Query" name:"DstCidrBlock"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteBgpNetworkRequest) Invoke(client goaliyun.Client) (*DeleteBgpNetworkResponse, error) {
	resp := &DeleteBgpNetworkResponse{}
	err := client.Request("vpc", "DeleteBgpNetwork", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBgpNetworkResponse struct {
	RequestId goaliyun.String
}
