package vpc

import (
	"github.com/morlay/goaliyun"
)

type EnableVpcClassicLinkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *EnableVpcClassicLinkRequest) Invoke(client goaliyun.Client) (*EnableVpcClassicLinkResponse, error) {
	resp := &EnableVpcClassicLinkResponse{}
	err := client.Request("vpc", "EnableVpcClassicLink", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnableVpcClassicLinkResponse struct {
	RequestId goaliyun.String
}
