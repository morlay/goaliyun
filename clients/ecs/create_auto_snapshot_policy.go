package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateAutoSnapshotPolicyRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	TimePoints             string `position:"Query" name:"TimePoints"`
	RetentionDays          int64  `position:"Query" name:"RetentionDays"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RepeatWeekdays         string `position:"Query" name:"RepeatWeekdays"`
	AutoSnapshotPolicyName string `position:"Query" name:"AutoSnapshotPolicyName"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *CreateAutoSnapshotPolicyRequest) Invoke(client goaliyun.Client) (*CreateAutoSnapshotPolicyResponse, error) {
	resp := &CreateAutoSnapshotPolicyResponse{}
	err := client.Request("ecs", "CreateAutoSnapshotPolicy", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAutoSnapshotPolicyResponse struct {
	RequestId            goaliyun.String
	AutoSnapshotPolicyId goaliyun.String
}
