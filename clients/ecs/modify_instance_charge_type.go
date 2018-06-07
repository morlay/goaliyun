package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceChargeTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	DryRun               string `position:"Query" name:"DryRun"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	IncludeDataDisks     string `position:"Query" name:"IncludeDataDisks"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeriodUnit           string `position:"Query" name:"PeriodUnit"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceChargeTypeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceChargeTypeResponse, error) {
	resp := &ModifyInstanceChargeTypeResponse{}
	err := client.Request("ecs", "ModifyInstanceChargeType", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceChargeTypeResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
