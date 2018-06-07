package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifySnapshotAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	SnapshotName         string `position:"Query" name:"SnapshotName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySnapshotAttributeRequest) Invoke(client goaliyun.Client) (*ModifySnapshotAttributeResponse, error) {
	resp := &ModifySnapshotAttributeResponse{}
	err := client.Request("ecs", "ModifySnapshotAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySnapshotAttributeResponse struct {
	RequestId goaliyun.String
}
