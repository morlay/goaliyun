package ecs

import (
	"github.com/morlay/goaliyun"
)

type DetachClassicLinkVpcRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DetachClassicLinkVpcRequest) Invoke(client goaliyun.Client) (*DetachClassicLinkVpcResponse, error) {
	resp := &DetachClassicLinkVpcResponse{}
	err := client.Request("ecs", "DetachClassicLinkVpc", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachClassicLinkVpcResponse struct {
	RequestId goaliyun.String
}
