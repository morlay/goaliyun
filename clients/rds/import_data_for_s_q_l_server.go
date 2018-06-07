package rds

import (
	"github.com/morlay/goaliyun"
)

type ImportDataForSQLServerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FileName             string `position:"Query" name:"FileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ImportDataForSQLServerRequest) Invoke(client goaliyun.Client) (*ImportDataForSQLServerResponse, error) {
	resp := &ImportDataForSQLServerResponse{}
	err := client.Request("rds", "ImportDataForSQLServer", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ImportDataForSQLServerResponse struct {
	RequestId goaliyun.String
	ImportID  goaliyun.Integer
}
