package ecs

import (
	"github.com/morlay/goaliyun"
)

type AddBandwidthPackageIpsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IpCount              string `position:"Query" name:"IpCount"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddBandwidthPackageIpsRequest) Invoke(client goaliyun.Client) (*AddBandwidthPackageIpsResponse, error) {
	resp := &AddBandwidthPackageIpsResponse{}
	err := client.Request("ecs", "AddBandwidthPackageIps", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddBandwidthPackageIpsResponse struct {
	RequestId goaliyun.String
}
