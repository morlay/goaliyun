package vpc

import (
	"github.com/morlay/goaliyun"
)

type UnassociateHaVipRequest struct {
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UnassociateHaVipRequest) Invoke(client goaliyun.Client) (*UnassociateHaVipResponse, error) {
	resp := &UnassociateHaVipResponse{}
	err := client.Request("vpc", "UnassociateHaVip", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnassociateHaVipResponse struct {
	RequestId goaliyun.String
}
