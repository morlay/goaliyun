package slb

import (
	"github.com/morlay/goaliyun"
)

type SetLoadBalancerAutoReleaseTimeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime      int64  `position:"Query" name:"AutoReleaseTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetLoadBalancerAutoReleaseTimeRequest) Invoke(client goaliyun.Client) (*SetLoadBalancerAutoReleaseTimeResponse, error) {
	resp := &SetLoadBalancerAutoReleaseTimeResponse{}
	err := client.Request("slb", "SetLoadBalancerAutoReleaseTime", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLoadBalancerAutoReleaseTimeResponse struct {
	RequestId goaliyun.String
}
