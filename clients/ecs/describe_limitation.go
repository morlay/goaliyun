package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeLimitationRequest struct {
	Limitation           string `position:"Query" name:"Limitation"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeLimitationRequest) Invoke(client goaliyun.Client) (*DescribeLimitationResponse, error) {
	resp := &DescribeLimitationResponse{}
	err := client.Request("ecs", "DescribeLimitation", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLimitationResponse struct {
	RequestId  goaliyun.String
	Limitation goaliyun.String
	Value      goaliyun.String
}
