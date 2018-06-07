package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateMigrateTaskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IsOnlineDB           string `position:"Query" name:"IsOnlineDB"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	OssObjectPositions   string `position:"Query" name:"OssObjectPositions"`
	OSSUrls              string `position:"Query" name:"OSSUrls"`
	DBName               string `position:"Query" name:"DBName"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BackupMode           string `position:"Query" name:"BackupMode"`
	CheckDBMode          string `position:"Query" name:"CheckDBMode"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateMigrateTaskRequest) Invoke(client goaliyun.Client) (*CreateMigrateTaskResponse, error) {
	resp := &CreateMigrateTaskResponse{}
	err := client.Request("rds", "CreateMigrateTask", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMigrateTaskResponse struct {
	RequestId     goaliyun.String
	DBInstanceId  goaliyun.String
	TaskId        goaliyun.String
	DBName        goaliyun.String
	MigrateTaskId goaliyun.String
	BackupMode    goaliyun.String
}
