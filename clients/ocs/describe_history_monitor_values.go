package ocs

import (
	"github.com/morlay/goaliyun"
)

type DescribeHistoryMonitorValuesRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	EndTime              string `position:"Query" name:"EndTime"`
	IntervalForHistory   string `position:"Query" name:"IntervalForHistory"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeHistoryMonitorValuesRequest) Invoke(client goaliyun.Client) (*DescribeHistoryMonitorValuesResponse, error) {
	resp := &DescribeHistoryMonitorValuesResponse{}
	err := client.Request("ocs", "DescribeHistoryMonitorValues", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeHistoryMonitorValuesResponse struct {
	RequestId      goaliyun.String
	MonitorHistory goaliyun.String
}
