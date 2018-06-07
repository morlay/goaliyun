package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAutoSnapshotPolicyRequest) Invoke(client goaliyun.Client) (*DescribeAutoSnapshotPolicyResponse, error) {
	resp := &DescribeAutoSnapshotPolicyResponse{}
	err := client.Request("ecs", "DescribeAutoSnapshotPolicy", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAutoSnapshotPolicyResponse struct {
	RequestId                  goaliyun.String
	AutoSnapshotOccupation     goaliyun.Integer
	AutoSnapshotPolicy         DescribeAutoSnapshotPolicyAutoSnapshotPolicy
	AutoSnapshotExcutionStatus DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus
}

type DescribeAutoSnapshotPolicyAutoSnapshotPolicy struct {
	SystemDiskPolicyEnabled           goaliyun.String
	SystemDiskPolicyTimePeriod        goaliyun.String
	SystemDiskPolicyRetentionDays     goaliyun.String
	SystemDiskPolicyRetentionLastWeek goaliyun.String
	DataDiskPolicyEnabled             goaliyun.String
	DataDiskPolicyTimePeriod          goaliyun.String
	DataDiskPolicyRetentionDays       goaliyun.String
	DataDiskPolicyRetentionLastWeek   goaliyun.String
}

type DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus struct {
	SystemDiskExcutionStatus goaliyun.String
	DataDiskExcutionStatus   goaliyun.String
}
