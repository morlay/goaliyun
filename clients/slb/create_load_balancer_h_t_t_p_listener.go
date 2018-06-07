package slb

import (
	"github.com/morlay/goaliyun"
)

type CreateLoadBalancerHTTPListenerRequest struct {
	Access_key_id          string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout     int64  `position:"Query" name:"HealthCheckTimeout"`
	ListenerForward        string `position:"Query" name:"ListenerForward"`
	XForwardedFor          string `position:"Query" name:"XForwardedFor"`
	HealthCheckURI         string `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold     int64  `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       int64  `position:"Query" name:"HealthyThreshold"`
	AclStatus              string `position:"Query" name:"AclStatus"`
	Scheduler              string `position:"Query" name:"Scheduler"`
	AclType                string `position:"Query" name:"AclType"`
	HealthCheck            string `position:"Query" name:"HealthCheck"`
	ForwardPort            int64  `position:"Query" name:"ForwardPort"`
	MaxConnection          int64  `position:"Query" name:"MaxConnection"`
	CookieTimeout          int64  `position:"Query" name:"CookieTimeout"`
	StickySessionType      string `position:"Query" name:"StickySessionType"`
	VpcIds                 string `position:"Query" name:"VpcIds"`
	VServerGroupId         string `position:"Query" name:"VServerGroupId"`
	AclId                  string `position:"Query" name:"AclId"`
	ListenerPort           int64  `position:"Query" name:"ListenerPort"`
	Cookie                 string `position:"Query" name:"Cookie"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth              int64  `position:"Query" name:"Bandwidth"`
	StickySession          string `position:"Query" name:"StickySession"`
	HealthCheckDomain      string `position:"Query" name:"HealthCheckDomain"`
	RequestTimeout         int64  `position:"Query" name:"RequestTimeout"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Gzip                   string `position:"Query" name:"Gzip"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	Tags                   string `position:"Query" name:"Tags"`
	IdleTimeout            int64  `position:"Query" name:"IdleTimeout"`
	LoadBalancerId         string `position:"Query" name:"LoadBalancerId"`
	XForwardedFor_SLBIP    string `position:"Query" name:"XForwardedFor_SLBIP"`
	BackendServerPort      int64  `position:"Query" name:"BackendServerPort"`
	HealthCheckInterval    int64  `position:"Query" name:"HealthCheckInterval"`
	XForwardedFor_proto    string `position:"Query" name:"XForwardedFor_proto"`
	XForwardedFor_SLBID    string `position:"Query" name:"XForwardedFor_SLBID"`
	HealthCheckConnectPort int64  `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string `position:"Query" name:"HealthCheckHttpCode"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *CreateLoadBalancerHTTPListenerRequest) Invoke(client goaliyun.Client) (*CreateLoadBalancerHTTPListenerResponse, error) {
	resp := &CreateLoadBalancerHTTPListenerResponse{}
	err := client.Request("slb", "CreateLoadBalancerHTTPListener", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateLoadBalancerHTTPListenerResponse struct {
	RequestId goaliyun.String
}
