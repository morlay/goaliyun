package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateTempDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RestoreTime          string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             int64  `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateTempDBInstanceRequest) Invoke(client goaliyun.Client) (*CreateTempDBInstanceResponse, error) {
	resp := &CreateTempDBInstanceResponse{}
	err := client.Request("rds", "CreateTempDBInstance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateTempDBInstanceResponse struct {
	RequestId        goaliyun.String
	TempDBInstanceId goaliyun.String
}
