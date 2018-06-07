package slb

import (
	"github.com/morlay/goaliyun"
)

type StartLoadBalancerListenerRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int64  `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *StartLoadBalancerListenerRequest) Invoke(client goaliyun.Client) (*StartLoadBalancerListenerResponse, error) {
	resp := &StartLoadBalancerListenerResponse{}
	err := client.Request("slb", "StartLoadBalancerListener", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartLoadBalancerListenerResponse struct {
	RequestId goaliyun.String
}
