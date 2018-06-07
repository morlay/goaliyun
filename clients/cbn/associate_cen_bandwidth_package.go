package cbn

import (
	"github.com/morlay/goaliyun"
)

type AssociateCenBandwidthPackageRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *AssociateCenBandwidthPackageRequest) Invoke(client goaliyun.Client) (*AssociateCenBandwidthPackageResponse, error) {
	resp := &AssociateCenBandwidthPackageResponse{}
	err := client.Request("cbn", "AssociateCenBandwidthPackage", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssociateCenBandwidthPackageResponse struct {
	RequestId goaliyun.String
}
