package rds

import (
	"github.com/morlay/goaliyun"
)

type DeleteDatabaseRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteDatabaseRequest) Invoke(client goaliyun.Client) (*DeleteDatabaseResponse, error) {
	resp := &DeleteDatabaseResponse{}
	err := client.Request("rds", "DeleteDatabase", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDatabaseResponse struct {
	RequestId goaliyun.String
}
