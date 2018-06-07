package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateHaVipRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateHaVipRequest) Invoke(client goaliyun.Client) (*CreateHaVipResponse, error) {
	resp := &CreateHaVipResponse{}
	err := client.Request("ecs", "CreateHaVip", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateHaVipResponse struct {
	RequestId goaliyun.String
	HaVipId   goaliyun.String
}
