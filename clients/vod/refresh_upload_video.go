package vod

import (
	"github.com/morlay/goaliyun"
)

type RefreshUploadVideoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RefreshUploadVideoRequest) Invoke(client goaliyun.Client) (*RefreshUploadVideoResponse, error) {
	resp := &RefreshUploadVideoResponse{}
	err := client.Request("vod", "RefreshUploadVideo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefreshUploadVideoResponse struct {
	RequestId     goaliyun.String
	UploadAuth    goaliyun.String
	UploadAddress goaliyun.String
}
