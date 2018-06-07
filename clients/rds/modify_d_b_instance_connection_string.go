package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceConnectionStringRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	Port                    string `position:"Query" name:"Port"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceConnectionStringRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceConnectionStringResponse, error) {
	resp := &ModifyDBInstanceConnectionStringResponse{}
	err := client.Request("rds", "ModifyDBInstanceConnectionString", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceConnectionStringResponse struct {
	RequestId goaliyun.String
}
