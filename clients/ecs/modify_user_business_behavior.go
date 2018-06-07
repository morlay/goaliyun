package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyUserBusinessBehaviorRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	StatusValue          string `position:"Query" name:"StatusValue"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	StatusKey            string `position:"Query" name:"StatusKey"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyUserBusinessBehaviorRequest) Invoke(client goaliyun.Client) (*ModifyUserBusinessBehaviorResponse, error) {
	resp := &ModifyUserBusinessBehaviorResponse{}
	err := client.Request("ecs", "ModifyUserBusinessBehavior", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyUserBusinessBehaviorResponse struct {
	RequestId goaliyun.String
}
