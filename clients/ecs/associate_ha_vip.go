package ecs

import (
	"github.com/morlay/goaliyun"
)

type AssociateHaVipRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AssociateHaVipRequest) Invoke(client goaliyun.Client) (*AssociateHaVipResponse, error) {
	resp := &AssociateHaVipResponse{}
	err := client.Request("ecs", "AssociateHaVip", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssociateHaVipResponse struct {
	RequestId goaliyun.String
}
