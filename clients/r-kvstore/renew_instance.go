package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type RenewInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceClass        string `position:"Query" name:"InstanceClass"`
	Capacity             string `position:"Query" name:"Capacity"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ForceUpgrade         string `position:"Query" name:"ForceUpgrade"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RenewInstanceRequest) Invoke(client goaliyun.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("r-kvstore", "RenewInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
	EndTime   goaliyun.String
}
