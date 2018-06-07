package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceNetworkExpireTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	ClassicExpiredDays   int64  `position:"Query" name:"ClassicExpiredDays"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceNetworkExpireTimeRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceNetworkExpireTimeResponse, error) {
	resp := &ModifyDBInstanceNetworkExpireTimeResponse{}
	err := client.Request("rds", "ModifyDBInstanceNetworkExpireTime", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceNetworkExpireTimeResponse struct {
	RequestId goaliyun.String
}
