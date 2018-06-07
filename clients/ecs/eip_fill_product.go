package ecs

import (
	"github.com/morlay/goaliyun"
)

type EipFillProductRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *EipFillProductRequest) Invoke(client goaliyun.Client) (*EipFillProductResponse, error) {
	resp := &EipFillProductResponse{}
	err := client.Request("ecs", "EipFillProduct", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EipFillProductResponse struct {
	RequestId goaliyun.String
	Data      goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
}
