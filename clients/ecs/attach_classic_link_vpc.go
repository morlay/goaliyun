package ecs

import (
	"github.com/morlay/goaliyun"
)

type AttachClassicLinkVpcRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AttachClassicLinkVpcRequest) Invoke(client goaliyun.Client) (*AttachClassicLinkVpcResponse, error) {
	resp := &AttachClassicLinkVpcResponse{}
	err := client.Request("ecs", "AttachClassicLinkVpc", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachClassicLinkVpcResponse struct {
	RequestId goaliyun.String
}
