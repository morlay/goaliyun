package rds

import (
	"github.com/morlay/goaliyun"
)

type UpgradeDBInstanceNetWorkInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpgradeDBInstanceNetWorkInfoRequest) Invoke(client goaliyun.Client) (*UpgradeDBInstanceNetWorkInfoResponse, error) {
	resp := &UpgradeDBInstanceNetWorkInfoResponse{}
	err := client.Request("rds", "UpgradeDBInstanceNetWorkInfo", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeDBInstanceNetWorkInfoResponse struct {
	RequestId goaliyun.String
}
