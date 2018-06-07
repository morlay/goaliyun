package polardb

import (
	"github.com/morlay/goaliyun"
)

type ModifyBackupPolicyRequest struct {
	PreferredBackupTime   string `position:"Query" name:"PreferredBackupTime"`
	PreferredBackupPeriod string `position:"Query" name:"PreferredBackupPeriod"`
	BackupRetentionPeriod string `position:"Query" name:"BackupRetentionPeriod"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId           string `position:"Query" name:"DBClusterId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ModifyBackupPolicyRequest) Invoke(client goaliyun.Client) (*ModifyBackupPolicyResponse, error) {
	resp := &ModifyBackupPolicyResponse{}
	err := client.Request("polardb", "ModifyBackupPolicy", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyBackupPolicyResponse struct {
	RequestId goaliyun.String
}
