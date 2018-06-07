package ecs

import (
	"github.com/morlay/goaliyun"
)

type ConnectRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ConnectRouterInterfaceRequest) Invoke(client goaliyun.Client) (*ConnectRouterInterfaceResponse, error) {
	resp := &ConnectRouterInterfaceResponse{}
	err := client.Request("ecs", "ConnectRouterInterface", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConnectRouterInterfaceResponse struct {
	RequestId goaliyun.String
}
