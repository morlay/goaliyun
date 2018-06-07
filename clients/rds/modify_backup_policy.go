package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyBackupPolicyRequest struct {
	PreferredBackupTime      string `position:"Query" name:"PreferredBackupTime"`
	PreferredBackupPeriod    string `position:"Query" name:"PreferredBackupPeriod"`
	BackupRetentionPeriod    string `position:"Query" name:"BackupRetentionPeriod"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	BackupLog                string `position:"Query" name:"BackupLog"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	LogBackupRetentionPeriod string `position:"Query" name:"LogBackupRetentionPeriod"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *ModifyBackupPolicyRequest) Invoke(client goaliyun.Client) (*ModifyBackupPolicyResponse, error) {
	resp := &ModifyBackupPolicyResponse{}
	err := client.Request("rds", "ModifyBackupPolicy", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyBackupPolicyResponse struct {
	RequestId goaliyun.String
}
