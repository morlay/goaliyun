package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DescribeHistoryMonitorValuesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	IntervalForHistory   string `position:"Query" name:"IntervalForHistory"`
	NodeId               string `position:"Query" name:"NodeId"`
	MonitorKeys          string `position:"Query" name:"MonitorKeys"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeHistoryMonitorValuesRequest) Invoke(client goaliyun.Client) (*DescribeHistoryMonitorValuesResponse, error) {
	resp := &DescribeHistoryMonitorValuesResponse{}
	err := client.Request("r-kvstore", "DescribeHistoryMonitorValues", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeHistoryMonitorValuesResponse struct {
	RequestId      goaliyun.String
	MonitorHistory goaliyun.String
}
