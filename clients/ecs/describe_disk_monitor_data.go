package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDiskMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DiskId               string `position:"Query" name:"DiskId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDiskMonitorDataRequest) Invoke(client goaliyun.Client) (*DescribeDiskMonitorDataResponse, error) {
	resp := &DescribeDiskMonitorDataResponse{}
	err := client.Request("ecs", "DescribeDiskMonitorData", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDiskMonitorDataResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	MonitorData DescribeDiskMonitorDataDiskMonitorDataList
}

type DescribeDiskMonitorDataDiskMonitorData struct {
	DiskId       goaliyun.String
	IOPSRead     goaliyun.Integer
	IOPSWrite    goaliyun.Integer
	IOPSTotal    goaliyun.Integer
	BPSRead      goaliyun.Integer
	BPSWrite     goaliyun.Integer
	BPSTotal     goaliyun.Integer
	LatencyRead  goaliyun.Integer
	LatencyWrite goaliyun.Integer
	TimeStamp    goaliyun.String
}

type DescribeDiskMonitorDataDiskMonitorDataList []DescribeDiskMonitorDataDiskMonitorData

func (list *DescribeDiskMonitorDataDiskMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDiskMonitorDataDiskMonitorData)
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
