package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyDiskAttributeRequest struct {
	DiskName             string `position:"Query" name:"DiskName"`
	DeleteAutoSnapshot   string `position:"Query" name:"DeleteAutoSnapshot"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	EnableAutoSnapshot   string `position:"Query" name:"EnableAutoSnapshot"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DeleteWithInstance   string `position:"Query" name:"DeleteWithInstance"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyDiskAttributeRequest) Invoke(client goaliyun.Client) (*ModifyDiskAttributeResponse, error) {
	resp := &ModifyDiskAttributeResponse{}
	err := client.Request("ecs", "ModifyDiskAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDiskAttributeResponse struct {
	RequestId goaliyun.String
}
