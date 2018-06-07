package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSnapshotMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSnapshotMonitorDataRequest) Invoke(client goaliyun.Client) (*DescribeSnapshotMonitorDataResponse, error) {
	resp := &DescribeSnapshotMonitorDataResponse{}
	err := client.Request("ecs", "DescribeSnapshotMonitorData", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSnapshotMonitorDataResponse struct {
	RequestId   goaliyun.String
	MonitorData DescribeSnapshotMonitorDataDataPointList
}

type DescribeSnapshotMonitorDataDataPoint struct {
	TimeStamp goaliyun.String
	Size      goaliyun.Integer
}

type DescribeSnapshotMonitorDataDataPointList []DescribeSnapshotMonitorDataDataPoint

func (list *DescribeSnapshotMonitorDataDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotMonitorDataDataPoint)
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
