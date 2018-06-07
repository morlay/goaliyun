package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteSnapshotRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteSnapshotRequest) Invoke(client goaliyun.Client) (*DeleteSnapshotResponse, error) {
	resp := &DeleteSnapshotResponse{}
	err := client.Request("ecs", "DeleteSnapshot", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSnapshotResponse struct {
	RequestId goaliyun.String
}
