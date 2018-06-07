package vpc

import (
	"github.com/morlay/goaliyun"
)

type TerminateVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *TerminateVirtualBorderRouterRequest) Invoke(client goaliyun.Client) (*TerminateVirtualBorderRouterResponse, error) {
	resp := &TerminateVirtualBorderRouterResponse{}
	err := client.Request("vpc", "TerminateVirtualBorderRouter", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TerminateVirtualBorderRouterResponse struct {
	RequestId goaliyun.String
}
