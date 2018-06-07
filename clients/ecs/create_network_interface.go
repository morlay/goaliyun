package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateNetworkInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Description          string `position:"Query" name:"Description"`
	NetworkInterfaceName string `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrimaryIpAddress     string `position:"Query" name:"PrimaryIpAddress"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateNetworkInterfaceRequest) Invoke(client goaliyun.Client) (*CreateNetworkInterfaceResponse, error) {
	resp := &CreateNetworkInterfaceResponse{}
	err := client.Request("ecs", "CreateNetworkInterface", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNetworkInterfaceResponse struct {
	RequestId          goaliyun.String
	NetworkInterfaceId goaliyun.String
}
