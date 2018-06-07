package ess

import (
	"github.com/morlay/goaliyun"
)

type EnableScalingGroupRequest struct {
	InstanceId10                 string `position:"Query" name:"InstanceId.10"`
	LoadBalancerWeight6          int64  `position:"Query" name:"LoadBalancerWeight.6"`
	LoadBalancerWeight11         int64  `position:"Query" name:"LoadBalancerWeight.11"`
	LoadBalancerWeight7          int64  `position:"Query" name:"LoadBalancerWeight.7"`
	LoadBalancerWeight12         int64  `position:"Query" name:"LoadBalancerWeight.12"`
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId12                 string `position:"Query" name:"InstanceId.12"`
	LoadBalancerWeight8          int64  `position:"Query" name:"LoadBalancerWeight.8"`
	InstanceId11                 string `position:"Query" name:"InstanceId.11"`
	LoadBalancerWeight9          int64  `position:"Query" name:"LoadBalancerWeight.9"`
	LoadBalancerWeight10         int64  `position:"Query" name:"LoadBalancerWeight.10"`
	LoadBalancerWeight2          int64  `position:"Query" name:"LoadBalancerWeight.2"`
	LoadBalancerWeight15         int64  `position:"Query" name:"LoadBalancerWeight.15"`
	LoadBalancerWeight3          int64  `position:"Query" name:"LoadBalancerWeight.3"`
	LoadBalancerWeight16         int64  `position:"Query" name:"LoadBalancerWeight.16"`
	ScalingGroupId               string `position:"Query" name:"ScalingGroupId"`
	LoadBalancerWeight4          int64  `position:"Query" name:"LoadBalancerWeight.4"`
	LoadBalancerWeight13         int64  `position:"Query" name:"LoadBalancerWeight.13"`
	LoadBalancerWeight5          int64  `position:"Query" name:"LoadBalancerWeight.5"`
	LoadBalancerWeight14         int64  `position:"Query" name:"LoadBalancerWeight.14"`
	ActiveScalingConfigurationId string `position:"Query" name:"ActiveScalingConfigurationId"`
	LoadBalancerWeight1          int64  `position:"Query" name:"LoadBalancerWeight.1"`
	InstanceId20                 string `position:"Query" name:"InstanceId.20"`
	InstanceId1                  string `position:"Query" name:"InstanceId.1"`
	LoadBalancerWeight20         int64  `position:"Query" name:"LoadBalancerWeight.20"`
	InstanceId3                  string `position:"Query" name:"InstanceId.3"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceId2                  string `position:"Query" name:"InstanceId.2"`
	InstanceId5                  string `position:"Query" name:"InstanceId.5"`
	InstanceId4                  string `position:"Query" name:"InstanceId.4"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	InstanceId7                  string `position:"Query" name:"InstanceId.7"`
	InstanceId6                  string `position:"Query" name:"InstanceId.6"`
	InstanceId9                  string `position:"Query" name:"InstanceId.9"`
	InstanceId8                  string `position:"Query" name:"InstanceId.8"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	InstanceId18                 string `position:"Query" name:"InstanceId.18"`
	LoadBalancerWeight19         int64  `position:"Query" name:"LoadBalancerWeight.19"`
	InstanceId17                 string `position:"Query" name:"InstanceId.17"`
	LoadBalancerWeight17         int64  `position:"Query" name:"LoadBalancerWeight.17"`
	InstanceId19                 string `position:"Query" name:"InstanceId.19"`
	LoadBalancerWeight18         int64  `position:"Query" name:"LoadBalancerWeight.18"`
	InstanceId14                 string `position:"Query" name:"InstanceId.14"`
	InstanceId13                 string `position:"Query" name:"InstanceId.13"`
	InstanceId16                 string `position:"Query" name:"InstanceId.16"`
	InstanceId15                 string `position:"Query" name:"InstanceId.15"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *EnableScalingGroupRequest) Invoke(client goaliyun.Client) (*EnableScalingGroupResponse, error) {
	resp := &EnableScalingGroupResponse{}
	err := client.Request("ess", "EnableScalingGroup", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnableScalingGroupResponse struct {
	RequestId goaliyun.String
}
