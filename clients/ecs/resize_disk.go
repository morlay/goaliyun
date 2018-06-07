package ecs

import (
	"github.com/morlay/goaliyun"
)

type ResizeDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NewSize              int64  `position:"Query" name:"NewSize"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ResizeDiskRequest) Invoke(client goaliyun.Client) (*ResizeDiskResponse, error) {
	resp := &ResizeDiskResponse{}
	err := client.Request("ecs", "ResizeDisk", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResizeDiskResponse struct {
	RequestId goaliyun.String
}
