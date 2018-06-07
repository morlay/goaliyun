package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeSnapshotsUsageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSnapshotsUsageRequest) Invoke(client goaliyun.Client) (*DescribeSnapshotsUsageResponse, error) {
	resp := &DescribeSnapshotsUsageResponse{}
	err := client.Request("ecs", "DescribeSnapshotsUsage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSnapshotsUsageResponse struct {
	RequestId     goaliyun.String
	SnapshotCount goaliyun.Integer
	SnapshotSize  goaliyun.Integer
}
