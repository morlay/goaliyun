package slb

import (
	"github.com/morlay/goaliyun"
)

type ModifyLoadBalancerInternetSpecRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            int64  `position:"Query" name:"Bandwidth"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyLoadBalancerInternetSpecRequest) Invoke(client goaliyun.Client) (*ModifyLoadBalancerInternetSpecResponse, error) {
	resp := &ModifyLoadBalancerInternetSpecResponse{}
	err := client.Request("slb", "ModifyLoadBalancerInternetSpec", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyLoadBalancerInternetSpecResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.Integer
}
