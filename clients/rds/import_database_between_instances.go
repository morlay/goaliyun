package rds

import (
	"github.com/morlay/goaliyun"
)

type ImportDatabaseBetweenInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SourceDBInstanceId   string `position:"Query" name:"SourceDBInstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInfo               string `position:"Query" name:"DBInfo"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ImportDatabaseBetweenInstancesRequest) Invoke(client goaliyun.Client) (*ImportDatabaseBetweenInstancesResponse, error) {
	resp := &ImportDatabaseBetweenInstancesResponse{}
	err := client.Request("rds", "ImportDatabaseBetweenInstances", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ImportDatabaseBetweenInstancesResponse struct {
	RequestId goaliyun.String
	ImportId  goaliyun.String
}
