package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceVncPasswdRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VncPassword          string `position:"Query" name:"VncPassword"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceVncPasswdRequest) Invoke(client goaliyun.Client) (*ModifyInstanceVncPasswdResponse, error) {
	resp := &ModifyInstanceVncPasswdResponse{}
	err := client.Request("ecs", "ModifyInstanceVncPasswd", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceVncPasswdResponse struct {
	RequestId goaliyun.String
}
