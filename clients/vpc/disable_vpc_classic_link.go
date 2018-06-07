package vpc

import (
	"github.com/morlay/goaliyun"
)

type DisableVpcClassicLinkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DisableVpcClassicLinkRequest) Invoke(client goaliyun.Client) (*DisableVpcClassicLinkResponse, error) {
	resp := &DisableVpcClassicLinkResponse{}
	err := client.Request("vpc", "DisableVpcClassicLink", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DisableVpcClassicLinkResponse struct {
	RequestId goaliyun.String
}
