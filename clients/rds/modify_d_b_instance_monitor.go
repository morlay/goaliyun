package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceMonitorRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               string `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceMonitorRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceMonitorResponse, error) {
	resp := &ModifyDBInstanceMonitorResponse{}
	err := client.Request("rds", "ModifyDBInstanceMonitor", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceMonitorResponse struct {
	RequestId goaliyun.String
}
