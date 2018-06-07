package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteRouterInterfaceRequest) Invoke(client goaliyun.Client) (*DeleteRouterInterfaceResponse, error) {
	resp := &DeleteRouterInterfaceResponse{}
	err := client.Request("ecs", "DeleteRouterInterface", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteRouterInterfaceResponse struct {
	RequestId goaliyun.String
}
