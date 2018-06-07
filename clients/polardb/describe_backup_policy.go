package polardb

import (
	"github.com/morlay/goaliyun"
)

type DescribeBackupPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBackupPolicyRequest) Invoke(client goaliyun.Client) (*DescribeBackupPolicyResponse, error) {
	resp := &DescribeBackupPolicyResponse{}
	err := client.Request("polardb", "DescribeBackupPolicy", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBackupPolicyResponse struct {
	RequestId               goaliyun.String
	BackupRetentionPeriod   goaliyun.Integer
	PreferredNextBackupTime goaliyun.String
	PreferredBackupTime     goaliyun.String
	PreferredBackupPeriod   goaliyun.String
}
