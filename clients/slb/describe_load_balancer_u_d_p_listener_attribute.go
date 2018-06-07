package slb

import (
	"github.com/morlay/goaliyun"
)

type DescribeLoadBalancerUDPListenerAttributeRequest struct {
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

func (req *DescribeLoadBalancerUDPListenerAttributeRequest) Invoke(client goaliyun.Client) (*DescribeLoadBalancerUDPListenerAttributeResponse, error) {
	resp := &DescribeLoadBalancerUDPListenerAttributeResponse{}
	err := client.Request("slb", "DescribeLoadBalancerUDPListenerAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	RequestId                 goaliyun.String
	ListenerPort              goaliyun.Integer
	BackendServerPort         goaliyun.Integer
	Status                    goaliyun.String
	Bandwidth                 goaliyun.Integer
	Scheduler                 goaliyun.String
	PersistenceTimeout        goaliyun.Integer
	HealthCheck               goaliyun.String
	HealthyThreshold          goaliyun.Integer
	UnhealthyThreshold        goaliyun.Integer
	HealthCheckConnectTimeout goaliyun.Integer
	HealthCheckConnectPort    goaliyun.Integer
	HealthCheckInterval       goaliyun.Integer
	HealthCheckReq            goaliyun.String
	HealthCheckExp            goaliyun.String
	MaxConnection             goaliyun.Integer
	VServerGroupId            goaliyun.String
	MasterSlaveServerGroupId  goaliyun.String
	AclId                     goaliyun.String
	AclType                   goaliyun.String
	AclStatus                 goaliyun.String
	VpcIds                    goaliyun.String
}
