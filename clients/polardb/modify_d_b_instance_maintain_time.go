package polardb

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceMaintainTimeRequest struct {
	MaintainTime         string `position:"Query" name:"MaintainTime"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceMaintainTimeRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceMaintainTimeResponse, error) {
	resp := &ModifyDBInstanceMaintainTimeResponse{}
	err := client.Request("polardb", "ModifyDBInstanceMaintainTime", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId goaliyun.String
}
