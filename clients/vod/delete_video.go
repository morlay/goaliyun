package vod

import (
	"github.com/morlay/goaliyun"
)

type DeleteVideoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VideoIds             string `position:"Query" name:"VideoIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteVideoRequest) Invoke(client goaliyun.Client) (*DeleteVideoResponse, error) {
	resp := &DeleteVideoResponse{}
	err := client.Request("vod", "DeleteVideo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVideoResponse struct {
	RequestId goaliyun.String
}
