package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSnapshotLinksRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	SnapshotLinkIds      string `position:"Query" name:"SnapshotLinkIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSnapshotLinksRequest) Invoke(client goaliyun.Client) (*DescribeSnapshotLinksResponse, error) {
	resp := &DescribeSnapshotLinksResponse{}
	err := client.Request("ecs", "DescribeSnapshotLinks", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSnapshotLinksResponse struct {
	RequestId     goaliyun.String
	TotalCount    goaliyun.Integer
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	SnapshotLinks DescribeSnapshotLinksSnapshotLinkList
}

type DescribeSnapshotLinksSnapshotLink struct {
	SnapshotLinkId goaliyun.String
	RegionId       goaliyun.String
	InstanceId     goaliyun.String
	InstanceName   goaliyun.String
	SourceDiskId   goaliyun.String
	SourceDiskSize goaliyun.Integer
	SourceDiskType goaliyun.String
	TotalSize      goaliyun.Integer
	TotalCount     goaliyun.Integer
}

type DescribeSnapshotLinksSnapshotLinkList []DescribeSnapshotLinksSnapshotLink

func (list *DescribeSnapshotLinksSnapshotLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotLinksSnapshotLink)
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
