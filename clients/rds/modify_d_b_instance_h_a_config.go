package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceHAConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SyncMode             string `position:"Query" name:"SyncMode"`
	DbInstanceId         string `position:"Query" name:"DbInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	HAMode               string `position:"Query" name:"HAMode"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceHAConfigRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceHAConfigResponse, error) {
	resp := &ModifyDBInstanceHAConfigResponse{}
	err := client.Request("rds", "ModifyDBInstanceHAConfig", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceHAConfigResponse struct {
	RequestId goaliyun.String
}
