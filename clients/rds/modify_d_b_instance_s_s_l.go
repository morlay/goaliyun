package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBInstanceSSLRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBInstanceSSLRequest) Invoke(client goaliyun.Client) (*ModifyDBInstanceSSLResponse, error) {
	resp := &ModifyDBInstanceSSLResponse{}
	err := client.Request("rds", "ModifyDBInstanceSSL", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBInstanceSSLResponse struct {
	RequestId goaliyun.String
}
