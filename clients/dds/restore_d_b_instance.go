package dds

import (
	"github.com/morlay/goaliyun"
)

type RestoreDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             int64  `position:"Query" name:"BackupId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RestoreDBInstanceRequest) Invoke(client goaliyun.Client) (*RestoreDBInstanceResponse, error) {
	resp := &RestoreDBInstanceResponse{}
	err := client.Request("dds", "RestoreDBInstance", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestoreDBInstanceResponse struct {
	RequestId goaliyun.String
}
