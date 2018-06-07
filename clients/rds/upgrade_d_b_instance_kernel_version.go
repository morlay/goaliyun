package rds

import (
	"github.com/morlay/goaliyun"
)

type UpgradeDBInstanceKernelVersionRequest struct {
	SwitchTimeMode       string `position:"Query" name:"SwitchTimeMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	SwitchTime           string `position:"Query" name:"SwitchTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpgradeDBInstanceKernelVersionRequest) Invoke(client goaliyun.Client) (*UpgradeDBInstanceKernelVersionResponse, error) {
	resp := &UpgradeDBInstanceKernelVersionResponse{}
	err := client.Request("rds", "UpgradeDBInstanceKernelVersion", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeDBInstanceKernelVersionResponse struct {
	RequestId          goaliyun.String
	DBInstanceName     goaliyun.String
	TaskId             goaliyun.String
	TargetMinorVersion goaliyun.String
}
