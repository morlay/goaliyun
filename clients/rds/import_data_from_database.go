package rds

import (
	"github.com/morlay/goaliyun"
)

type ImportDataFromDatabaseRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ImportDataType         string `position:"Query" name:"ImportDataType"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	IsLockTable            string `position:"Query" name:"IsLockTable"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	SourceDatabaseDBNames  string `position:"Query" name:"SourceDatabaseDBNames"`
	SourceDatabaseIp       string `position:"Query" name:"SourceDatabaseIp"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	SourceDatabasePassword string `position:"Query" name:"SourceDatabasePassword"`
	SourceDatabasePort     string `position:"Query" name:"SourceDatabasePort"`
	SourceDatabaseUserName string `position:"Query" name:"SourceDatabaseUserName"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *ImportDataFromDatabaseRequest) Invoke(client goaliyun.Client) (*ImportDataFromDatabaseResponse, error) {
	resp := &ImportDataFromDatabaseResponse{}
	err := client.Request("rds", "ImportDataFromDatabase", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ImportDataFromDatabaseResponse struct {
	RequestId goaliyun.String
	ImportId  goaliyun.Integer
}
