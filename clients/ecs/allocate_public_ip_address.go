package ecs

import (
	"github.com/morlay/goaliyun"
)

type AllocatePublicIpAddressRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VlanId               string `position:"Query" name:"VlanId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AllocatePublicIpAddressRequest) Invoke(client goaliyun.Client) (*AllocatePublicIpAddressResponse, error) {
	resp := &AllocatePublicIpAddressResponse{}
	err := client.Request("ecs", "AllocatePublicIpAddress", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AllocatePublicIpAddressResponse struct {
	RequestId goaliyun.String
	IpAddress goaliyun.String
}
