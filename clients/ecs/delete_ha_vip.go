package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteHaVipRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteHaVipRequest) Invoke(client goaliyun.Client) (*DeleteHaVipResponse, error) {
	resp := &DeleteHaVipResponse{}
	err := client.Request("ecs", "DeleteHaVip", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteHaVipResponse struct {
	RequestId goaliyun.String
}
