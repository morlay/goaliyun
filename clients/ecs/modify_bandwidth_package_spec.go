package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyBandwidthPackageSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyBandwidthPackageSpecRequest) Invoke(client goaliyun.Client) (*ModifyBandwidthPackageSpecResponse, error) {
	resp := &ModifyBandwidthPackageSpecResponse{}
	err := client.Request("ecs", "ModifyBandwidthPackageSpec", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyBandwidthPackageSpecResponse struct {
	RequestId goaliyun.String
}
