package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeClusterOperationHostTaskLogRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	TaskId          string `position:"Query" name:"TaskId"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterOperationHostTaskLogRequest) Invoke(client goaliyun.Client) (*DescribeClusterOperationHostTaskLogResponse, error) {
	resp := &DescribeClusterOperationHostTaskLogResponse{}
	err := client.Request("emr", "DescribeClusterOperationHostTaskLog", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterOperationHostTaskLogResponse struct {
	RequestId goaliyun.String
	Stdout    goaliyun.String
	Stderr    goaliyun.String
}
