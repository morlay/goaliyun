package ess

import (
	"github.com/morlay/goaliyun"
)

type DescribeScalingActivityDetailRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingActivityId    string `position:"Query" name:"ScalingActivityId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeScalingActivityDetailRequest) Invoke(client goaliyun.Client) (*DescribeScalingActivityDetailResponse, error) {
	resp := &DescribeScalingActivityDetailResponse{}
	err := client.Request("ess", "DescribeScalingActivityDetail", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeScalingActivityDetailResponse struct {
	ScalingActivityId goaliyun.String
	Detail            goaliyun.String
}
