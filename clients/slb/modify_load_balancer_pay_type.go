package slb

import (
	"github.com/morlay/goaliyun"
)

type ModifyLoadBalancerPayTypeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	Duration             int64  `position:"Query" name:"Duration"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	PayType              string `position:"Query" name:"PayType"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyLoadBalancerPayTypeRequest) Invoke(client goaliyun.Client) (*ModifyLoadBalancerPayTypeResponse, error) {
	resp := &ModifyLoadBalancerPayTypeResponse{}
	err := client.Request("slb", "ModifyLoadBalancerPayType", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyLoadBalancerPayTypeResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.Integer
}
