package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteImageRequest) Invoke(client goaliyun.Client) (*DeleteImageResponse, error) {
	resp := &DeleteImageResponse{}
	err := client.Request("ecs", "DeleteImage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteImageResponse struct {
	RequestId goaliyun.String
}
