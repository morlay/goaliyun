package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyBackupPolicyRequest struct {
	PreferredBackupTime   string `position:"Query" name:"PreferredBackupTime"`
	PreferredBackupPeriod string `position:"Query" name:"PreferredBackupPeriod"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId            string `position:"Query" name:"InstanceId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ModifyBackupPolicyRequest) Invoke(client goaliyun.Client) (*ModifyBackupPolicyResponse, error) {
	resp := &ModifyBackupPolicyResponse{}
	err := client.Request("r-kvstore", "ModifyBackupPolicy", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyBackupPolicyResponse struct {
	RequestId goaliyun.String
}
