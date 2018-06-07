package cbn

import (
	"github.com/morlay/goaliyun"
)

type SetCenInterRegionBandwidthLimitRequest struct {
	LocalRegionId        string `position:"Query" name:"LocalRegionId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OppositeRegionId     string `position:"Query" name:"OppositeRegionId"`
	BandwidthLimit       int64  `position:"Query" name:"BandwidthLimit"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetCenInterRegionBandwidthLimitRequest) Invoke(client goaliyun.Client) (*SetCenInterRegionBandwidthLimitResponse, error) {
	resp := &SetCenInterRegionBandwidthLimitResponse{}
	err := client.Request("cbn", "SetCenInterRegionBandwidthLimit", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetCenInterRegionBandwidthLimitResponse struct {
	RequestId goaliyun.String
}
