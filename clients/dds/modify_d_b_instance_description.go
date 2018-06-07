package dds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceDescriptionRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	DBInstanceId          string `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	NodeId                string `position:"Query" name:"NodeId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceDescriptionResponse, error) {
	resp := &ModifyDBInstanceDescriptionResponse{}
	err := client.Request("dds", "ModifyDBInstanceDescription", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceDescriptionResponse struct {
	RequestId goaliyun.String
}
