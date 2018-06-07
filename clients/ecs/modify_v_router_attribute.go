package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyVRouterAttributeRequest struct {
	VRouterName          string `position:"Query" name:"VRouterName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVRouterAttributeRequest) Invoke(client goaliyun.Client) (*ModifyVRouterAttributeResponse, error) {
	resp := &ModifyVRouterAttributeResponse{}
	err := client.Request("ecs", "ModifyVRouterAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVRouterAttributeResponse struct {
	RequestId goaliyun.String
}
