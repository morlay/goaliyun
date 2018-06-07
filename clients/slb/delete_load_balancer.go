package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteLoadBalancerRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteLoadBalancerRequest) Invoke(client goaliyun.Client) (*DeleteLoadBalancerResponse, error) {
	resp := &DeleteLoadBalancerResponse{}
	err := client.Request("slb", "DeleteLoadBalancer", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteLoadBalancerResponse struct {
	RequestId goaliyun.String
}
