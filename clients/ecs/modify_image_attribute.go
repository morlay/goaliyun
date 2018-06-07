package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyImageAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ImageName            string `position:"Query" name:"ImageName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyImageAttributeRequest) Invoke(client goaliyun.Client) (*ModifyImageAttributeResponse, error) {
	resp := &ModifyImageAttributeResponse{}
	err := client.Request("ecs", "ModifyImageAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyImageAttributeResponse struct {
	RequestId goaliyun.String
}
