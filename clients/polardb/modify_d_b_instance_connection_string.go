package polardb

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceConnectionStringRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceConnectionStringRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceConnectionStringResponse, error) {
	resp := &ModifyDBInstanceConnectionStringResponse{}
	err := client.Request("polardb", "ModifyDBInstanceConnectionString", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceConnectionStringResponse struct {
	RequestId           goaliyun.String
	DBInstanceId        goaliyun.String
	OldConnectionString goaliyun.String
	OldPort             goaliyun.String
	NewConnectionString goaliyun.String
	NewPort             goaliyun.String
	IPType              goaliyun.String
}
