package ecs

import (
	"github.com/morlay/goaliyun"
)

type RenewInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	PeriodUnit           string `position:"Query" name:"PeriodUnit"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RenewInstanceRequest) Invoke(client goaliyun.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("ecs", "RenewInstance", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	RequestId goaliyun.String
}
