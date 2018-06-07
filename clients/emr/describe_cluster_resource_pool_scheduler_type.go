package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeClusterResourcePoolSchedulerTypeRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterResourcePoolSchedulerTypeRequest) Invoke(client goaliyun.Client) (*DescribeClusterResourcePoolSchedulerTypeResponse, error) {
	resp := &DescribeClusterResourcePoolSchedulerTypeResponse{}
	err := client.Request("emr", "DescribeClusterResourcePoolSchedulerType", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterResourcePoolSchedulerTypeResponse struct {
	RequestId            goaliyun.String
	CurrentSchedulerType goaliyun.String
	SupportSchedulerType goaliyun.String
}
