package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyParameterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Forcerestart         string `position:"Query" name:"Forcerestart"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Parameters           string `position:"Query" name:"Parameters"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyParameterRequest) Invoke(client goaliyun.Client) (*ModifyParameterResponse, error) {
	resp := &ModifyParameterResponse{}
	err := client.Request("rds", "ModifyParameter", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyParameterResponse struct {
	RequestId goaliyun.String
}
