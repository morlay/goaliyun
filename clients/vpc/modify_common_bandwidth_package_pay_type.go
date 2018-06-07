package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyCommonBandwidthPackagePayTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Duration             int64  `position:"Query" name:"Duration"`
	KbpsBandwidth        string `position:"Query" name:"KbpsBandwidth"`
	ResourceUid          int64  `position:"Query" name:"ResourceUid"`
	ResourceBid          string `position:"Query" name:"ResourceBid"`
	PayType              string `position:"Query" name:"PayType"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyCommonBandwidthPackagePayTypeRequest) Invoke(client goaliyun.Client) (*ModifyCommonBandwidthPackagePayTypeResponse, error) {
	resp := &ModifyCommonBandwidthPackagePayTypeResponse{}
	err := client.Request("vpc", "ModifyCommonBandwidthPackagePayType", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCommonBandwidthPackagePayTypeResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.Integer
	Code      goaliyun.String
	Message   goaliyun.String
}
