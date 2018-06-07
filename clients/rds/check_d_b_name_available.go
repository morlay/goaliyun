package rds

import (
	"github.com/morlay/goaliyun"
)

type CheckDBNameAvailableRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CheckDBNameAvailableRequest) Invoke(client goaliyun.Client) (*CheckDBNameAvailableResponse, error) {
	resp := &CheckDBNameAvailableResponse{}
	err := client.Request("rds", "CheckDBNameAvailable", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckDBNameAvailableResponse struct {
	RequestId goaliyun.String
}
