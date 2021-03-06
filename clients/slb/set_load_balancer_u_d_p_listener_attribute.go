package slb

import (
	"github.com/morlay/goaliyun"
)

type SetLoadBalancerUDPListenerAttributeRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	HealthCheckConnectTimeout int64  `position:"Query" name:"HealthCheckConnectTimeout"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	UnhealthyThreshold        int64  `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          int64  `position:"Query" name:"HealthyThreshold"`
	AclStatus                 string `position:"Query" name:"AclStatus"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	AclType                   string `position:"Query" name:"AclType"`
	MasterSlaveServerGroup    string `position:"Query" name:"MasterSlaveServerGroup"`
	MaxConnection             int64  `position:"Query" name:"MaxConnection"`
	PersistenceTimeout        int64  `position:"Query" name:"PersistenceTimeout"`
	VpcIds                    string `position:"Query" name:"VpcIds"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	AclId                     string `position:"Query" name:"AclId"`
	ListenerPort              int64  `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 int64  `position:"Query" name:"Bandwidth"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckReq            string `position:"Query" name:"HealthCheckReq"`
	HealthCheckInterval       int64  `position:"Query" name:"HealthCheckInterval"`
	HealthCheckExp            string `position:"Query" name:"HealthCheckExp"`
	HealthCheckConnectPort    int64  `position:"Query" name:"HealthCheckConnectPort"`
	VServerGroup              string `position:"Query" name:"VServerGroup"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *SetLoadBalancerUDPListenerAttributeRequest) Invoke(client goaliyun.Client) (*SetLoadBalancerUDPListenerAttributeResponse, error) {
	resp := &SetLoadBalancerUDPListenerAttributeResponse{}
	err := client.Request("slb", "SetLoadBalancerUDPListenerAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLoadBalancerUDPListenerAttributeResponse struct {
	RequestId goaliyun.String
}
