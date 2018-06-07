package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyHaVipAttributeRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyHaVipAttributeRequest) Invoke(client goaliyun.Client) (*ModifyHaVipAttributeResponse, error) {
	resp := &ModifyHaVipAttributeResponse{}
	err := client.Request("ecs", "ModifyHaVipAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyHaVipAttributeResponse struct {
	RequestId goaliyun.String
}
