package rds

import (
	"github.com/morlay/goaliyun"
)

type RenewInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               string `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RenewInstanceRequest) Invoke(client goaliyun.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("rds", "RenewInstance", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.Integer
}
