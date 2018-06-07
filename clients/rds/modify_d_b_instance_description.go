package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceDescriptionRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceDescriptionResponse, error) {
	resp := &ModifyDBInstanceDescriptionResponse{}
	err := client.Request("rds", "ModifyDBInstanceDescription", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceDescriptionResponse struct {
	RequestId goaliyun.String
}
