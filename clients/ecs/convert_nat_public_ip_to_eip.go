package ecs

import (
	"github.com/morlay/goaliyun"
)

type ConvertNatPublicIpToEipRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ConvertNatPublicIpToEipRequest) Invoke(client goaliyun.Client) (*ConvertNatPublicIpToEipResponse, error) {
	resp := &ConvertNatPublicIpToEipResponse{}
	err := client.Request("ecs", "ConvertNatPublicIpToEip", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConvertNatPublicIpToEipResponse struct {
	RequestId goaliyun.String
}
