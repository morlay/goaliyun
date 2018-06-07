package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAutoSnapshotPolicyExRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeAutoSnapshotPolicyExRequest) Invoke(client goaliyun.Client) (*DescribeAutoSnapshotPolicyExResponse, error) {
	resp := &DescribeAutoSnapshotPolicyExResponse{}
	err := client.Request("ecs", "DescribeAutoSnapshotPolicyEx", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAutoSnapshotPolicyExResponse struct {
	RequestId            goaliyun.String
	TotalCount           goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	AutoSnapshotPolicies DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicy struct {
	AutoSnapshotPolicyId   goaliyun.String
	RegionId               goaliyun.String
	AutoSnapshotPolicyName goaliyun.String
	TimePoints             goaliyun.String
	RepeatWeekdays         goaliyun.String
	RetentionDays          goaliyun.Integer
	DiskNums               goaliyun.Integer
	VolumeNums             goaliyun.Integer
	CreationTime           goaliyun.String
	Status                 goaliyun.String
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList []DescribeAutoSnapshotPolicyExAutoSnapshotPolicy

func (list *DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAutoSnapshotPolicyExAutoSnapshotPolicy)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
