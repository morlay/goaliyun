package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeClusterOperationHostTaskLogForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	TaskId          string `position:"Query" name:"TaskId"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterOperationHostTaskLogForAdminRequest) Invoke(client goaliyun.Client) (*DescribeClusterOperationHostTaskLogForAdminResponse, error) {
	resp := &DescribeClusterOperationHostTaskLogForAdminResponse{}
	err := client.Request("emr", "DescribeClusterOperationHostTaskLogForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterOperationHostTaskLogForAdminResponse struct {
	RequestId goaliyun.String
	Stdout    goaliyun.String
	Stderr    goaliyun.String
}
