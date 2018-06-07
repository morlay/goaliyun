package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceClass        string `position:"Query" name:"InstanceClass"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ForceUpgrade         string `position:"Query" name:"ForceUpgrade"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyInstanceSpecResponse, error) {
	resp := &ModifyInstanceSpecResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceSpec", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceSpecResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
