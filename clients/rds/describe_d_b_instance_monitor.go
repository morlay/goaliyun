package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceMonitorRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceMonitorRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceMonitorResponse, error) {
	resp := &DescribeDBInstanceMonitorResponse{}
	err := client.Request("rds", "DescribeDBInstanceMonitor", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceMonitorResponse struct {
	RequestId goaliyun.String
	Period    goaliyun.String
}
