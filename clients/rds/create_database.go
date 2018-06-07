package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateDatabaseRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	DBDescription        string `position:"Query" name:"DBDescription"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CharacterSetName     string `position:"Query" name:"CharacterSetName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateDatabaseRequest) Invoke(client goaliyun.Client) (*CreateDatabaseResponse, error) {
	resp := &CreateDatabaseResponse{}
	err := client.Request("rds", "CreateDatabase", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDatabaseResponse struct {
	RequestId goaliyun.String
}
