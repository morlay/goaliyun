package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceAutoRenewAttributeRequest struct {
	Duration             int64  `position:"Query" name:"Duration"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PeriodUnit           string `position:"Query" name:"PeriodUnit"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoRenew            string `position:"Query" name:"AutoRenew"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RenewalStatus        string `position:"Query" name:"RenewalStatus"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceAutoRenewAttributeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceAutoRenewAttributeResponse, error) {
	resp := &ModifyInstanceAutoRenewAttributeResponse{}
	err := client.Request("ecs", "ModifyInstanceAutoRenewAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceAutoRenewAttributeResponse struct {
	RequestId goaliyun.String
}
