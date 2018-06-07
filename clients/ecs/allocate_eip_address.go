package ecs

import (
	"github.com/morlay/goaliyun"
)

type AllocateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AllocateEipAddressRequest) Invoke(client goaliyun.Client) (*AllocateEipAddressResponse, error) {
	resp := &AllocateEipAddressResponse{}
	err := client.Request("ecs", "AllocateEipAddress", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AllocateEipAddressResponse struct {
	RequestId    goaliyun.String
	AllocationId goaliyun.String
	EipAddress   goaliyun.String
}
