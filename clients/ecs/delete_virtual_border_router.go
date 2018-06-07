package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteVirtualBorderRouterRequest) Invoke(client goaliyun.Client) (*DeleteVirtualBorderRouterResponse, error) {
	resp := &DeleteVirtualBorderRouterResponse{}
	err := client.Request("ecs", "DeleteVirtualBorderRouter", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVirtualBorderRouterResponse struct {
	RequestId goaliyun.String
}
