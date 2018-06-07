package ecs

import (
	"github.com/morlay/goaliyun"
)

type ReleasePublicIpAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PublicIpAddress      string `position:"Query" name:"PublicIpAddress"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReleasePublicIpAddressRequest) Invoke(client goaliyun.Client) (*ReleasePublicIpAddressResponse, error) {
	resp := &ReleasePublicIpAddressResponse{}
	err := client.Request("ecs", "ReleasePublicIpAddress", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleasePublicIpAddressResponse struct {
	RequestId goaliyun.String
}
