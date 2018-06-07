package r_kvstore

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSnapshotsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	BeginTime            string `position:"Query" name:"BeginTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSnapshotsRequest) Invoke(client goaliyun.Client) (*DescribeSnapshotsResponse, error) {
	resp := &DescribeSnapshotsResponse{}
	err := client.Request("r-kvstore", "DescribeSnapshots", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSnapshotsResponse struct {
	RequestId goaliyun.String
	Snapshots DescribeSnapshotsSnapshotList
}

type DescribeSnapshotsSnapshot struct {
	SnapshotId         goaliyun.String
	SnapshotName       goaliyun.String
	InstanceId         goaliyun.String
	CreateTime         goaliyun.String
	Memory             goaliyun.Integer
	RdbSize            goaliyun.Integer
	Status             goaliyun.String
	Type               goaliyun.String
	OssDownloadInPath  goaliyun.String
	OssDownloadOutPath goaliyun.String
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
