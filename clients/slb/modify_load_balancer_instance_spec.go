package slb

import (
	"github.com/morlay/goaliyun"
)

type ModifyLoadBalancerInstanceSpecRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	LoadBalancerSpec     string `position:"Query" name:"LoadBalancerSpec"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyLoadBalancerInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyLoadBalancerInstanceSpecResponse, error) {
	resp := &ModifyLoadBalancerInstanceSpecResponse{}
	err := client.Request("slb", "ModifyLoadBalancerInstanceSpec", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyLoadBalancerInstanceSpecResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.Integer
}
