package rds

import (
	"github.com/morlay/goaliyun"
)

type UpgradeDBInstanceNetworkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpgradeDBInstanceNetworkRequest) Invoke(client goaliyun.Client) (*UpgradeDBInstanceNetworkResponse, error) {
	resp := &UpgradeDBInstanceNetworkResponse{}
	err := client.Request("rds", "UpgradeDBInstanceNetwork", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeDBInstanceNetworkResponse struct {
	RequestId      goaliyun.String
	DBInstanceName goaliyun.String
}
