package vpc

import (
	"github.com/morlay/goaliyun"
)

type AllocateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ISP                  string `position:"Query" name:"ISP"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	Netmode              string `position:"Query" name:"Netmode"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AllocateEipAddressRequest) Invoke(client goaliyun.Client) (*AllocateEipAddressResponse, error) {
	resp := &AllocateEipAddressResponse{}
	err := client.Request("vpc", "AllocateEipAddress", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AllocateEipAddressResponse struct {
	RequestId       goaliyun.String
	AllocationId    goaliyun.String
	EipAddress      goaliyun.String
	OrderId         goaliyun.Integer
	ResourceGroupId goaliyun.String
}
