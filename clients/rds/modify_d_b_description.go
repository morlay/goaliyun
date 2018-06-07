package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDBDescriptionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	DBDescription        string `position:"Query" name:"DBDescription"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDBDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyDBDescriptionResponse, error) {
	resp := &ModifyDBDescriptionResponse{}
	err := client.Request("rds", "ModifyDBDescription", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDBDescriptionResponse struct {
	RequestId goaliyun.String
}
