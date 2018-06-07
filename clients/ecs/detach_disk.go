package ecs

import (
	"github.com/morlay/goaliyun"
)

type DetachDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DetachDiskRequest) Invoke(client goaliyun.Client) (*DetachDiskResponse, error) {
	resp := &DetachDiskResponse{}
	err := client.Request("ecs", "DetachDisk", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachDiskResponse struct {
	RequestId goaliyun.String
}
