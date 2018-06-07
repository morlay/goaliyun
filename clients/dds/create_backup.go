package dds

import (
	"github.com/morlay/goaliyun"
)

type CreateBackupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateBackupRequest) Invoke(client goaliyun.Client) (*CreateBackupResponse, error) {
	resp := &CreateBackupResponse{}
	err := client.Request("dds", "CreateBackup", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBackupResponse struct {
	RequestId goaliyun.String
	BackupId  goaliyun.String
}
