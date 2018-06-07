package cbn

import (
	"github.com/morlay/goaliyun"
)

type UnassociateCenBandwidthPackageRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *UnassociateCenBandwidthPackageRequest) Invoke(client goaliyun.Client) (*UnassociateCenBandwidthPackageResponse, error) {
	resp := &UnassociateCenBandwidthPackageResponse{}
	err := client.Request("cbn", "UnassociateCenBandwidthPackage", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnassociateCenBandwidthPackageResponse struct {
	RequestId goaliyun.String
}
