package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyImageShareGroupPermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	AddGroup1            string `position:"Query" name:"AddGroup.1"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RemoveGroup1         string `position:"Query" name:"RemoveGroup.1"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyImageShareGroupPermissionRequest) Invoke(client goaliyun.Client) (*ModifyImageShareGroupPermissionResponse, error) {
	resp := &ModifyImageShareGroupPermissionResponse{}
	err := client.Request("ecs", "ModifyImageShareGroupPermission", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyImageShareGroupPermissionResponse struct {
	RequestId goaliyun.String
}
