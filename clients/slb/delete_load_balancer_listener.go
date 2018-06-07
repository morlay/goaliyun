package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteLoadBalancerListenerRequest struct {
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

func (req *DeleteLoadBalancerListenerRequest) Invoke(client goaliyun.Client) (*DeleteLoadBalancerListenerResponse, error) {
	resp := &DeleteLoadBalancerListenerResponse{}
	err := client.Request("slb", "DeleteLoadBalancerListener", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLoadBalancerListenerResponse struct {
	RequestId goaliyun.String
}
