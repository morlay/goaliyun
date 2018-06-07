package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyAutoSnapshotPolicyRequest struct {
	DataDiskPolicyEnabled             string `position:"Query" name:"DataDiskPolicyEnabled"`
	ResourceOwnerId                   int64  `position:"Query" name:"ResourceOwnerId"`
	DataDiskPolicyRetentionDays       int64  `position:"Query" name:"DataDiskPolicyRetentionDays"`
	ResourceOwnerAccount              string `position:"Query" name:"ResourceOwnerAccount"`
	SystemDiskPolicyRetentionLastWeek string `position:"Query" name:"SystemDiskPolicyRetentionLastWeek"`
	OwnerAccount                      string `position:"Query" name:"OwnerAccount"`
	SystemDiskPolicyTimePeriod        int64  `position:"Query" name:"SystemDiskPolicyTimePeriod"`
	OwnerId                           int64  `position:"Query" name:"OwnerId"`
	DataDiskPolicyRetentionLastWeek   string `position:"Query" name:"DataDiskPolicyRetentionLastWeek"`
	SystemDiskPolicyRetentionDays     int64  `position:"Query" name:"SystemDiskPolicyRetentionDays"`
	DataDiskPolicyTimePeriod          int64  `position:"Query" name:"DataDiskPolicyTimePeriod"`
	SystemDiskPolicyEnabled           string `position:"Query" name:"SystemDiskPolicyEnabled"`
	RegionId                          string `position:"Query" name:"RegionId"`
}

func (req *ModifyAutoSnapshotPolicyRequest) Invoke(client goaliyun.Client) (*ModifyAutoSnapshotPolicyResponse, error) {
	resp := &ModifyAutoSnapshotPolicyResponse{}
	err := client.Request("ecs", "ModifyAutoSnapshotPolicy", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAutoSnapshotPolicyResponse struct {
	RequestId goaliyun.String
}
