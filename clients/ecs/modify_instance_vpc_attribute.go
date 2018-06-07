package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceVpcAttributeRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceVpcAttributeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceVpcAttributeResponse, error) {
	resp := &ModifyInstanceVpcAttributeResponse{}
	err := client.Request("ecs", "ModifyInstanceVpcAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceVpcAttributeResponse struct {
	RequestId goaliyun.String
}
