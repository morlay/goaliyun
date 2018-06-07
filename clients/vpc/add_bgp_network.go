package vpc

import (
	"github.com/morlay/goaliyun"
)

type AddBgpNetworkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DstCidrBlock         string `position:"Query" name:"DstCidrBlock"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddBgpNetworkRequest) Invoke(client goaliyun.Client) (*AddBgpNetworkResponse, error) {
	resp := &AddBgpNetworkResponse{}
	err := client.Request("vpc", "AddBgpNetwork", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddBgpNetworkResponse struct {
	RequestId goaliyun.String
}
