package dds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceNetExpireTimeRequest struct {
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken            string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString         string `position:"Query" name:"ConnectionString"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	ClassicExpendExpiredDays int64  `position:"Query" name:"ClassicExpendExpiredDays"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceNetExpireTimeRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceNetExpireTimeResponse, error) {
	resp := &ModifyDBInstanceNetExpireTimeResponse{}
	err := client.Request("dds", "ModifyDBInstanceNetExpireTime", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceNetExpireTimeResponse struct {
	RequestId goaliyun.String
}
