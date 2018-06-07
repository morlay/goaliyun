package rds

import (
	"github.com/morlay/goaliyun"
)

type MigrateToOtherZoneRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	EffectiveTime        string `position:"Query" name:"EffectiveTime"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *MigrateToOtherZoneRequest) Invoke(client goaliyun.Client) (*MigrateToOtherZoneResponse, error) {
	resp := &MigrateToOtherZoneResponse{}
	err := client.Request("rds", "MigrateToOtherZone", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MigrateToOtherZoneResponse struct {
	RequestId goaliyun.String
}
