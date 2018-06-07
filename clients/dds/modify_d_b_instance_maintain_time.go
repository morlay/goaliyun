package dds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceMaintainTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	MaintainStartTime    string `position:"Query" name:"MaintainStartTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MaintainEndTime      string `position:"Query" name:"MaintainEndTime"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceMaintainTimeRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceMaintainTimeResponse, error) {
	resp := &ModifyDBInstanceMaintainTimeResponse{}
	err := client.Request("dds", "ModifyDBInstanceMaintainTime", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId goaliyun.String
}
