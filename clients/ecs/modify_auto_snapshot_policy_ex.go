package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyAutoSnapshotPolicyExRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId   string `position:"Query" name:"AutoSnapshotPolicyId"`
	TimePoints             string `position:"Query" name:"TimePoints"`
	RetentionDays          int64  `position:"Query" name:"RetentionDays"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RepeatWeekdays         string `position:"Query" name:"RepeatWeekdays"`
	AutoSnapshotPolicyName string `position:"Query" name:"AutoSnapshotPolicyName"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *ModifyAutoSnapshotPolicyExRequest) Invoke(client goaliyun.Client) (*ModifyAutoSnapshotPolicyExResponse, error) {
	resp := &ModifyAutoSnapshotPolicyExResponse{}
	err := client.Request("ecs", "ModifyAutoSnapshotPolicyEx", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAutoSnapshotPolicyExResponse struct {
	RequestId goaliyun.String
}
