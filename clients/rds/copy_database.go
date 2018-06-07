package rds

import (
	"github.com/morlay/goaliyun"
)

type CopyDatabaseRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CopyDatabaseRequest) Invoke(client goaliyun.Client) (*CopyDatabaseResponse, error) {
	resp := &CopyDatabaseResponse{}
	err := client.Request("rds", "CopyDatabase", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CopyDatabaseResponse struct {
	DBName   goaliyun.String
	DBStatus goaliyun.String
	TaskId   goaliyun.String
}
