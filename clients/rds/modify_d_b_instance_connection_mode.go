package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceConnectionModeRequest struct {
	ConnectionMode       string `position:"Query" name:"ConnectionMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceConnectionModeRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceConnectionModeResponse, error) {
	resp := &ModifyDBInstanceConnectionModeResponse{}
	err := client.Request("rds", "ModifyDBInstanceConnectionMode", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceConnectionModeResponse struct {
	RequestId goaliyun.String
}
