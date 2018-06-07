package cbn

import (
	"github.com/morlay/goaliyun"
)

type ModifyCenBandwidthPackageSpecRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth             int64  `position:"Query" name:"Bandwidth"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ModifyCenBandwidthPackageSpecRequest) Invoke(client goaliyun.Client) (*ModifyCenBandwidthPackageSpecResponse, error) {
	resp := &ModifyCenBandwidthPackageSpecResponse{}
	err := client.Request("cbn", "ModifyCenBandwidthPackageSpec", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCenBandwidthPackageSpecResponse struct {
	RequestId goaliyun.String
}
