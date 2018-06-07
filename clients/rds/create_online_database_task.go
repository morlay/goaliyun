package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateOnlineDatabaseTaskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string `position:"Query" name:"MigrateTaskId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	CheckDBMode          string `position:"Query" name:"CheckDBMode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateOnlineDatabaseTaskRequest) Invoke(client goaliyun.Client) (*CreateOnlineDatabaseTaskResponse, error) {
	resp := &CreateOnlineDatabaseTaskResponse{}
	err := client.Request("rds", "CreateOnlineDatabaseTask", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateOnlineDatabaseTaskResponse struct {
	RequestId goaliyun.String
}
