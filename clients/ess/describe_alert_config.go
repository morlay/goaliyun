package ess

import (
	"github.com/morlay/goaliyun"
)

type DescribeAlertConfigRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAlertConfigRequest) Invoke(client goaliyun.Client) (*DescribeAlertConfigResponse, error) {
	resp := &DescribeAlertConfigResponse{}
	err := client.Request("ess", "DescribeAlertConfig", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAlertConfigResponse struct {
	SuccessConfig goaliyun.Integer
	FailConfig    goaliyun.Integer
	RejectConfig  goaliyun.Integer
	RequestId     goaliyun.String
}
