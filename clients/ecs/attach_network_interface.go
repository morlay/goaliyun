package ecs

import (
	"github.com/morlay/goaliyun"
)

type AttachNetworkInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	NetworkInterfaceId   string `position:"Query" name:"NetworkInterfaceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AttachNetworkInterfaceRequest) Invoke(client goaliyun.Client) (*AttachNetworkInterfaceResponse, error) {
	resp := &AttachNetworkInterfaceResponse{}
	err := client.Request("ecs", "AttachNetworkInterface", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachNetworkInterfaceResponse struct {
	RequestId goaliyun.String
}
