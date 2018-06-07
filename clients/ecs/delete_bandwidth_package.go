package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteBandwidthPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteBandwidthPackageRequest) Invoke(client goaliyun.Client) (*DeleteBandwidthPackageResponse, error) {
	resp := &DeleteBandwidthPackageResponse{}
	err := client.Request("ecs", "DeleteBandwidthPackage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBandwidthPackageResponse struct {
	RequestId goaliyun.String
}
