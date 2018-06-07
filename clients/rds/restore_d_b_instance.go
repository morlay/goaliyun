package rds

import (
	"github.com/morlay/goaliyun"
)

type RestoreDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RestoreDBInstanceRequest) Invoke(client goaliyun.Client) (*RestoreDBInstanceResponse, error) {
	resp := &RestoreDBInstanceResponse{}
	err := client.Request("rds", "RestoreDBInstance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestoreDBInstanceResponse struct {
	RequestId goaliyun.String
}
