package ecs

import (
	"github.com/morlay/goaliyun"
)

type ResetDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ResetDiskRequest) Invoke(client goaliyun.Client) (*ResetDiskResponse, error) {
	resp := &ResetDiskResponse{}
	err := client.Request("ecs", "ResetDisk", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetDiskResponse struct {
	RequestId goaliyun.String
}
