package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteAutoSnapshotPolicyRequest) Invoke(client goaliyun.Client) (*DeleteAutoSnapshotPolicyResponse, error) {
	resp := &DeleteAutoSnapshotPolicyResponse{}
	err := client.Request("ecs", "DeleteAutoSnapshotPolicy", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAutoSnapshotPolicyResponse struct {
	RequestId goaliyun.String
}
