package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int64  `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PayType              string `position:"Query" name:"PayType"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceSpecResponse, error) {
	resp := &ModifyDBInstanceSpecResponse{}
	err := client.Request("rds", "ModifyDBInstanceSpec", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceSpecResponse struct {
	RequestId goaliyun.String
}
