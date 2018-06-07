package cbn

import (
	"github.com/morlay/goaliyun"
)

type CreateCenBandwidthPackageRequest struct {
	GeographicRegionBId        string `position:"Query" name:"GeographicRegionBId"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	Period                     int64  `position:"Query" name:"Period"`
	GeographicRegionAId        string `position:"Query" name:"GeographicRegionAId"`
	AutoPay                    string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                string `position:"Query" name:"ClientToken"`
	Bandwidth                  int64  `position:"Query" name:"Bandwidth"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	Description                string `position:"Query" name:"Description"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	BandwidthPackageChargeType string `position:"Query" name:"BandwidthPackageChargeType"`
	Name                       string `position:"Query" name:"Name"`
	PricingCycle               string `position:"Query" name:"PricingCycle"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *CreateCenBandwidthPackageRequest) Invoke(client goaliyun.Client) (*CreateCenBandwidthPackageResponse, error) {
	resp := &CreateCenBandwidthPackageResponse{}
	err := client.Request("cbn", "CreateCenBandwidthPackage", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateCenBandwidthPackageResponse struct {
	RequestId                  goaliyun.String
	CenBandwidthPackageId      goaliyun.String
	CenBandwidthPackageOrderId goaliyun.String
}
