package ecs

import (
	"github.com/morlay/goaliyun"
)

type ApplyAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ApplyAutoSnapshotPolicyRequest) Invoke(client goaliyun.Client) (*ApplyAutoSnapshotPolicyResponse, error) {
	resp := &ApplyAutoSnapshotPolicyResponse{}
	err := client.Request("ecs", "ApplyAutoSnapshotPolicy", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ApplyAutoSnapshotPolicyResponse struct {
	RequestId goaliyun.String
}
