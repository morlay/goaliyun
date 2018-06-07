package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateBackupRequest struct {
	BackupMethod         string `position:"Query" name:"BackupMethod"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackupType           string `position:"Query" name:"BackupType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateBackupRequest) Invoke(client goaliyun.Client) (*CreateBackupResponse, error) {
	resp := &CreateBackupResponse{}
	err := client.Request("rds", "CreateBackup", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBackupResponse struct {
	RequestId goaliyun.String
}
