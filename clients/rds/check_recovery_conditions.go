package rds

import (
	"github.com/morlay/goaliyun"
)

type CheckRecoveryConditionsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RestoreTime          string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupFile           string `position:"Query" name:"BackupFile"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CheckRecoveryConditionsRequest) Invoke(client goaliyun.Client) (*CheckRecoveryConditionsResponse, error) {
	resp := &CheckRecoveryConditionsResponse{}
	err := client.Request("rds", "CheckRecoveryConditions", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckRecoveryConditionsResponse struct {
	RequestId      goaliyun.String
	DBInstanceId   goaliyun.String
	RecoveryStatus goaliyun.String
}
