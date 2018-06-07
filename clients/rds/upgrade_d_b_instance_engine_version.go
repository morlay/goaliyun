package rds

import (
	"github.com/morlay/goaliyun"
)

type UpgradeDBInstanceEngineVersionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpgradeDBInstanceEngineVersionRequest) Invoke(client goaliyun.Client) (*UpgradeDBInstanceEngineVersionResponse, error) {
	resp := &UpgradeDBInstanceEngineVersionResponse{}
	err := client.Request("rds", "UpgradeDBInstanceEngineVersion", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeDBInstanceEngineVersionResponse struct {
	RequestId goaliyun.String
	TaskId    goaliyun.String
}
