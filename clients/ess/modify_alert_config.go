package ess

import (
	"github.com/morlay/goaliyun"
)

type ModifyAlertConfigRequest struct {
	SuccessConfig        int64  `position:"Query" name:"SuccessConfig"`
	RejectConfig         int64  `position:"Query" name:"RejectConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	FailConfig           int64  `position:"Query" name:"FailConfig"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyAlertConfigRequest) Invoke(client goaliyun.Client) (*ModifyAlertConfigResponse, error) {
	resp := &ModifyAlertConfigResponse{}
	err := client.Request("ess", "ModifyAlertConfig", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAlertConfigResponse struct {
	RequestId goaliyun.String
}
