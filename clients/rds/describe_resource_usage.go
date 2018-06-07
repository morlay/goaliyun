package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeResourceUsageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeResourceUsageRequest) Invoke(client goaliyun.Client) (*DescribeResourceUsageResponse, error) {
	resp := &DescribeResourceUsageResponse{}
	err := client.Request("rds", "DescribeResourceUsage", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeResourceUsageResponse struct {
	RequestId      goaliyun.String
	DBInstanceId   goaliyun.String
	Engine         goaliyun.String
	DiskUsed       goaliyun.Integer
	DataSize       goaliyun.Integer
	LogSize        goaliyun.Integer
	BackupSize     goaliyun.Integer
	SQLSize        goaliyun.Integer
	ColdBackupSize goaliyun.Integer
}
