package slb

import (
	"github.com/morlay/goaliyun"
)

type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
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

func (req *DescribeLoadBalancerHTTPSListenerAttributeRequest) Invoke(client goaliyun.Client) (*DescribeLoadBalancerHTTPSListenerAttributeResponse, error) {
	resp := &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	err := client.Request("slb", "DescribeLoadBalancerHTTPSListenerAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	RequestId              goaliyun.String
	ListenerPort           goaliyun.Integer
	BackendServerPort      goaliyun.Integer
	Bandwidth              goaliyun.Integer
	Status                 goaliyun.String
	SecurityStatus         goaliyun.String
	XForwardedFor          goaliyun.String
	Scheduler              goaliyun.String
	StickySession          goaliyun.String
	StickySessionType      goaliyun.String
	CookieTimeout          goaliyun.Integer
	Cookie                 goaliyun.String
	HealthCheck            goaliyun.String
	HealthCheckDomain      goaliyun.String
	HealthCheckURI         goaliyun.String
	HealthyThreshold       goaliyun.Integer
	UnhealthyThreshold     goaliyun.Integer
	HealthCheckTimeout     goaliyun.Integer
	HealthCheckInterval    goaliyun.Integer
	HealthCheckConnectPort goaliyun.Integer
	HealthCheckHttpCode    goaliyun.String
	ServerCertificateId    goaliyun.String
	CACertificateId        goaliyun.String
	MaxConnection          goaliyun.Integer
	VServerGroupId         goaliyun.String
	Gzip                   goaliyun.String
	XForwardedFor_SLBIP    goaliyun.String
	XForwardedFor_SLBID    goaliyun.String
	XForwardedFor_proto    goaliyun.String
	AclId                  goaliyun.String
	AclType                goaliyun.String
	AclStatus              goaliyun.String
	VpcIds                 goaliyun.String
	RequestTimeout         goaliyun.Integer
	IdleTimeout            goaliyun.Integer
}
