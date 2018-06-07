package ecs

import (
	"github.com/morlay/goaliyun"
)

type CancelCopyImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelCopyImageRequest) Invoke(client goaliyun.Client) (*CancelCopyImageResponse, error) {
	resp := &CancelCopyImageResponse{}
	err := client.Request("ecs", "CancelCopyImage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelCopyImageResponse struct {
	RequestId goaliyun.String
}
