package slb

import (
	"github.com/morlay/goaliyun"
)

type SetLoadBalancerNameRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerName     string `position:"Query" name:"LoadBalancerName"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetLoadBalancerNameRequest) Invoke(client goaliyun.Client) (*SetLoadBalancerNameResponse, error) {
	resp := &SetLoadBalancerNameResponse{}
	err := client.Request("slb", "SetLoadBalancerName", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLoadBalancerNameResponse struct {
	RequestId goaliyun.String
}
