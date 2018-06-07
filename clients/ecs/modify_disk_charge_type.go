package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyDiskChargeTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDiskChargeTypeRequest) Invoke(client goaliyun.Client) (*ModifyDiskChargeTypeResponse, error) {
	resp := &ModifyDiskChargeTypeResponse{}
	err := client.Request("ecs", "ModifyDiskChargeType", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDiskChargeTypeResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
