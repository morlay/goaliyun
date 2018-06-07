package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteDiskRequest) Invoke(client goaliyun.Client) (*DeleteDiskResponse, error) {
	resp := &DeleteDiskResponse{}
	err := client.Request("ecs", "DeleteDisk", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDiskResponse struct {
	RequestId goaliyun.String
}
