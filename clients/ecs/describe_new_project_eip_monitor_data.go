package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeNewProjectEipMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeNewProjectEipMonitorDataRequest) Invoke(client goaliyun.Client) (*DescribeNewProjectEipMonitorDataResponse, error) {
	resp := &DescribeNewProjectEipMonitorDataResponse{}
	err := client.Request("ecs", "DescribeNewProjectEipMonitorData", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNewProjectEipMonitorDataResponse struct {
	RequestId       goaliyun.String
	EipMonitorDatas DescribeNewProjectEipMonitorDataEipMonitorDataList
}

type DescribeNewProjectEipMonitorDataEipMonitorData struct {
	EipRX        goaliyun.Integer
	EipTX        goaliyun.Integer
	EipFlow      goaliyun.Integer
	EipBandwidth goaliyun.Integer
	EipPackets   goaliyun.Integer
	TimeStamp    goaliyun.String
}

type DescribeNewProjectEipMonitorDataEipMonitorDataList []DescribeNewProjectEipMonitorDataEipMonitorData

func (list *DescribeNewProjectEipMonitorDataEipMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNewProjectEipMonitorDataEipMonitorData)
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
