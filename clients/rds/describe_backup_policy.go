package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeBackupPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBackupPolicyRequest) Invoke(client goaliyun.Client) (*DescribeBackupPolicyResponse, error) {
	resp := &DescribeBackupPolicyResponse{}
	err := client.Request("rds", "DescribeBackupPolicy", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBackupPolicyResponse struct {
	RequestId                goaliyun.String
	BackupRetentionPeriod    goaliyun.Integer
	PreferredNextBackupTime  goaliyun.String
	PreferredBackupTime      goaliyun.String
	PreferredBackupPeriod    goaliyun.String
	BackupLog                goaliyun.String
	LogBackupRetentionPeriod goaliyun.Integer
}
