package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyPrepayInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	OperatorType         string `position:"Query" name:"OperatorType"`
	SystemDiskCategory   string `position:"Query" name:"SystemDiskCategory"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	MigrateAcrossZone    string `position:"Query" name:"MigrateAcrossZone"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyPrepayInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyPrepayInstanceSpecResponse, error) {
	resp := &ModifyPrepayInstanceSpecResponse{}
	err := client.Request("ecs", "ModifyPrepayInstanceSpec", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyPrepayInstanceSpecResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
