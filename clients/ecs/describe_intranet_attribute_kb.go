package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeIntranetAttributeKbRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeIntranetAttributeKbRequest) Invoke(client goaliyun.Client) (*DescribeIntranetAttributeKbResponse, error) {
	resp := &DescribeIntranetAttributeKbResponse{}
	err := client.Request("ecs", "DescribeIntranetAttributeKb", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeIntranetAttributeKbResponse struct {
	RequestId               goaliyun.String
	InstanceId              goaliyun.String
	VlanId                  goaliyun.String
	IntranetIpAddress       goaliyun.String
	IntranetMaxBandwidthIn  goaliyun.Integer
	IntranetMaxBandwidthOut goaliyun.Integer
}
