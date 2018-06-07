package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSnapshotsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Filter2Value         string `position:"Query" name:"Filter.2.Value"`
	SnapshotIds          string `position:"Query" name:"SnapshotIds"`
	Usage                string `position:"Query" name:"Usage"`
	SnapshotLinkId       string `position:"Query" name:"SnapshotLinkId"`
	SnapshotName         string `position:"Query" name:"SnapshotName"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Filter1Key           string `position:"Query" name:"Filter.1.Key"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DiskId               string `position:"Query" name:"DiskId"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	DryRun               string `position:"Query" name:"DryRun"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SourceDiskType       string `position:"Query" name:"SourceDiskType"`
	Filter1Value         string `position:"Query" name:"Filter.1.Value"`
	Filter2Key           string `position:"Query" name:"Filter.2.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	Encrypted            string `position:"Query" name:"Encrypted"`
	SnapshotType         string `position:"Query" name:"SnapshotType"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSnapshotsRequest) Invoke(client goaliyun.Client) (*DescribeSnapshotsResponse, error) {
	resp := &DescribeSnapshotsResponse{}
	err := client.Request("ecs", "DescribeSnapshots", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSnapshotsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Snapshots  DescribeSnapshotsSnapshotList
}

type DescribeSnapshotsSnapshot struct {
	SnapshotId        goaliyun.String
	SnapshotName      goaliyun.String
	Progress          goaliyun.String
	ProductCode       goaliyun.String
	SourceDiskId      goaliyun.String
	SourceDiskType    goaliyun.String
	RetentionDays     goaliyun.Integer
	Encrypted         bool
	SourceDiskSize    goaliyun.String
	Description       goaliyun.String
	CreationTime      goaliyun.String
	Status            goaliyun.String
	Usage             goaliyun.String
	SourceStorageType goaliyun.String
	Tags              DescribeSnapshotsTagList
}

type DescribeSnapshotsTag struct {
	TagKey   goaliyun.String
	TagValue goaliyun.String
}

type DescribeSnapshotsSnapshotList []DescribeSnapshotsSnapshot

func (list *DescribeSnapshotsSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotsSnapshot)
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

type DescribeSnapshotsTagList []DescribeSnapshotsTag

func (list *DescribeSnapshotsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotsTag)
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
