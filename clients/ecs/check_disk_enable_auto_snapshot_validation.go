package ecs

import (
	"github.com/morlay/goaliyun"
)

type CheckDiskEnableAutoSnapshotValidationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CheckDiskEnableAutoSnapshotValidationRequest) Invoke(client goaliyun.Client) (*CheckDiskEnableAutoSnapshotValidationResponse, error) {
	resp := &CheckDiskEnableAutoSnapshotValidationResponse{}
	err := client.Request("ecs", "CheckDiskEnableAutoSnapshotValidation", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckDiskEnableAutoSnapshotValidationResponse struct {
	RequestId              goaliyun.String
	IsPermitted            goaliyun.String
	AutoSnapshotOccupation goaliyun.Integer
}
