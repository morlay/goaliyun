package slb

import (
	"github.com/morlay/goaliyun"
)

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	Access_key_id          string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout     int64  `position:"Query" name:"HealthCheckTimeout"`
	XForwardedFor          string `position:"Query" name:"XForwardedFor"`
	HealthCheckURI         string `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold     int64  `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       int64  `position:"Query" name:"HealthyThreshold"`
	AclStatus              string `position:"Query" name:"AclStatus"`
	Scheduler              string `position:"Query" name:"Scheduler"`
	AclType                string `position:"Query" name:"AclType"`
	HealthCheck            string `position:"Query" name:"HealthCheck"`
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
	HealthCheckInterval    int64  `position:"Query" name:"HealthCheckInterval"`
	XForwardedFor_proto    string `position:"Query" name:"XForwardedFor_proto"`
	XForwardedFor_SLBID    string `position:"Query" name:"XForwardedFor_SLBID"`
	HealthCheckConnectPort int64  `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string `position:"Query" name:"HealthCheckHttpCode"`
	VServerGroup           string `position:"Query" name:"VServerGroup"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *SetLoadBalancerHTTPListenerAttributeRequest) Invoke(client goaliyun.Client) (*SetLoadBalancerHTTPListenerAttributeResponse, error) {
	resp := &SetLoadBalancerHTTPListenerAttributeResponse{}
	err := client.Request("slb", "SetLoadBalancerHTTPListenerAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLoadBalancerHTTPListenerAttributeResponse struct {
	RequestId goaliyun.String
}
