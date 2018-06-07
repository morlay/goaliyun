package cbn

import (
	"github.com/morlay/goaliyun"
)

type ModifyCenBandwidthPackageAttributeRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	Name                  string `position:"Query" name:"Name"`
	Description           string `position:"Query" name:"Description"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ModifyCenBandwidthPackageAttributeRequest) Invoke(client goaliyun.Client) (*ModifyCenBandwidthPackageAttributeResponse, error) {
	resp := &ModifyCenBandwidthPackageAttributeResponse{}
	err := client.Request("cbn", "ModifyCenBandwidthPackageAttribute", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCenBandwidthPackageAttributeResponse struct {
	RequestId goaliyun.String
}
