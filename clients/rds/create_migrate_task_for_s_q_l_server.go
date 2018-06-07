package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateMigrateTaskForSQLServerRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TaskType             string `position:"Query" name:"TaskType"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IsOnlineDB           string `position:"Query" name:"IsOnlineDB"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	OSSUrls              string `position:"Query" name:"OSSUrls"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateMigrateTaskForSQLServerRequest) Invoke(client goaliyun.Client) (*CreateMigrateTaskForSQLServerResponse, error) {
	resp := &CreateMigrateTaskForSQLServerResponse{}
	err := client.Request("rds", "CreateMigrateTaskForSQLServer", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMigrateTaskForSQLServerResponse struct {
	RequestId      goaliyun.String
	DBInstanceId   goaliyun.String
	DBInstanceName goaliyun.String
	TaskId         goaliyun.String
	DBName         goaliyun.String
	MigrateIaskId  goaliyun.String
	TaskType       goaliyun.String
}
