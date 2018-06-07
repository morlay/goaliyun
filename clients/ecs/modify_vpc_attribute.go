package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyVpcAttributeRequest struct {
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVpcAttributeRequest) Invoke(client goaliyun.Client) (*ModifyVpcAttributeResponse, error) {
	resp := &ModifyVpcAttributeResponse{}
	err := client.Request("ecs", "ModifyVpcAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVpcAttributeResponse struct {
	RequestId goaliyun.String
}
