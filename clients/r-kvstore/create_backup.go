package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type CreateBackupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateBackupRequest) Invoke(client goaliyun.Client) (*CreateBackupResponse, error) {
	resp := &CreateBackupResponse{}
	err := client.Request("r-kvstore", "CreateBackup", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBackupResponse struct {
	RequestId   goaliyun.String
	BackupJobID goaliyun.String
}
