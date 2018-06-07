package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceTDERequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TDEStatus            string `position:"Query" name:"TDEStatus"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceTDERequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceTDEResponse, error) {
	resp := &ModifyDBInstanceTDEResponse{}
	err := client.Request("rds", "ModifyDBInstanceTDE", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceTDEResponse struct {
	RequestId goaliyun.String
}
