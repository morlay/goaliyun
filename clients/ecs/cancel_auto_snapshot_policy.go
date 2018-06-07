package ecs

import (
	"github.com/morlay/goaliyun"
)

type CancelAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelAutoSnapshotPolicyRequest) Invoke(client goaliyun.Client) (*CancelAutoSnapshotPolicyResponse, error) {
	resp := &CancelAutoSnapshotPolicyResponse{}
	err := client.Request("ecs", "CancelAutoSnapshotPolicy", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelAutoSnapshotPolicyResponse struct {
	RequestId goaliyun.String
}
