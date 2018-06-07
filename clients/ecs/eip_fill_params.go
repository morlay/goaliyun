package ecs

import (
	"github.com/morlay/goaliyun"
)

type EipFillParamsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *EipFillParamsRequest) Invoke(client goaliyun.Client) (*EipFillParamsResponse, error) {
	resp := &EipFillParamsResponse{}
	err := client.Request("ecs", "EipFillParams", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EipFillParamsResponse struct {
	RequestId goaliyun.String
	Data      goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
}
