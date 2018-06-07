package polardb

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBClusterConnectionStringRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId             string `position:"Query" name:"DBClusterId"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBClusterConnectionStringRequest) Invoke(client goaliyun.Client) (*ModifyDBClusterConnectionStringResponse, error) {
	resp := &ModifyDBClusterConnectionStringResponse{}
	err := client.Request("polardb", "ModifyDBClusterConnectionString", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBClusterConnectionStringResponse struct {
	RequestId           goaliyun.String
	DBClusterId         goaliyun.String
	OldConnectionString goaliyun.String
	OldPort             goaliyun.String
	NewConnectionString goaliyun.String
	NewPort             goaliyun.String
	IPType              goaliyun.String
}
