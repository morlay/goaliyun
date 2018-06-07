package ecs

import (
	"github.com/morlay/goaliyun"
)

type AttachDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Device               string `position:"Query" name:"Device"`
	DeleteWithInstance   string `position:"Query" name:"DeleteWithInstance"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AttachDiskRequest) Invoke(client goaliyun.Client) (*AttachDiskResponse, error) {
	resp := &AttachDiskResponse{}
	err := client.Request("ecs", "AttachDisk", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachDiskResponse struct {
	RequestId goaliyun.String
}
