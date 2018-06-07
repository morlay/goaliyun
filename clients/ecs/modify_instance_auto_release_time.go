package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceAutoReleaseTimeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime      string `position:"Query" name:"AutoReleaseTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceAutoReleaseTimeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceAutoReleaseTimeResponse, error) {
	resp := &ModifyInstanceAutoReleaseTimeResponse{}
	err := client.Request("ecs", "ModifyInstanceAutoReleaseTime", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceAutoReleaseTimeResponse struct {
	RequestId goaliyun.String
}
