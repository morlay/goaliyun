package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type RenewMultiInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RenewMultiInstanceRequest) Invoke(client goaliyun.Client) (*RenewMultiInstanceResponse, error) {
	resp := &RenewMultiInstanceResponse{}
	err := client.Request("r-kvstore", "RenewMultiInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewMultiInstanceResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
