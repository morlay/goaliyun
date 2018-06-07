package cbn

import (
	"github.com/morlay/goaliyun"
)

type DeleteCenBandwidthPackageRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *DeleteCenBandwidthPackageRequest) Invoke(client goaliyun.Client) (*DeleteCenBandwidthPackageResponse, error) {
	resp := &DeleteCenBandwidthPackageResponse{}
	err := client.Request("cbn", "DeleteCenBandwidthPackage", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCenBandwidthPackageResponse struct {
	RequestId goaliyun.String
}
