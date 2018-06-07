package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteNetworkInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteNetworkInterfaceRequest) Invoke(client goaliyun.Client) (*DeleteNetworkInterfaceResponse, error) {
	resp := &DeleteNetworkInterfaceResponse{}
	err := client.Request("ecs", "DeleteNetworkInterface", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNetworkInterfaceResponse struct {
	RequestId goaliyun.String
}
