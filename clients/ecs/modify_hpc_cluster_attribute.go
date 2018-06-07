package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyHpcClusterAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	HpcClusterId         string `position:"Query" name:"HpcClusterId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Description          string `position:"Query" name:"Description"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyHpcClusterAttributeRequest) Invoke(client goaliyun.Client) (*ModifyHpcClusterAttributeResponse, error) {
	resp := &ModifyHpcClusterAttributeResponse{}
	err := client.Request("ecs", "ModifyHpcClusterAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyHpcClusterAttributeResponse struct {
	RequestId goaliyun.String
}
