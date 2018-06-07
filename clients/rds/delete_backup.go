package rds

import (
	"github.com/morlay/goaliyun"
)

type DeleteBackupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteBackupRequest) Invoke(client goaliyun.Client) (*DeleteBackupResponse, error) {
	resp := &DeleteBackupResponse{}
	err := client.Request("rds", "DeleteBackup", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteBackupResponse struct {
	RequestId goaliyun.String
}
