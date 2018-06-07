package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DescribeBackupPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBackupPolicyRequest) Invoke(client goaliyun.Client) (*DescribeBackupPolicyResponse, error) {
	resp := &DescribeBackupPolicyResponse{}
	err := client.Request("r-kvstore", "DescribeBackupPolicy", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBackupPolicyResponse struct {
	RequestId               goaliyun.String
	BackupRetentionPeriod   goaliyun.String
	PreferredBackupTime     goaliyun.String
	PreferredBackupPeriod   goaliyun.String
	PreferredNextBackupTime goaliyun.String
}
