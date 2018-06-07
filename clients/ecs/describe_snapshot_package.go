package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSnapshotPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSnapshotPackageRequest) Invoke(client goaliyun.Client) (*DescribeSnapshotPackageResponse, error) {
	resp := &DescribeSnapshotPackageResponse{}
	err := client.Request("ecs", "DescribeSnapshotPackage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSnapshotPackageResponse struct {
	RequestId        goaliyun.String
	TotalCount       goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageSize         goaliyun.Integer
	SnapshotPackages DescribeSnapshotPackageSnapshotPackageList
}

type DescribeSnapshotPackageSnapshotPackage struct {
	StartTime    goaliyun.String
	EndTime      goaliyun.String
	InitCapacity goaliyun.Integer
	DisplayName  goaliyun.String
}

type DescribeSnapshotPackageSnapshotPackageList []DescribeSnapshotPackageSnapshotPackage

func (list *DescribeSnapshotPackageSnapshotPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotPackageSnapshotPackage)
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
