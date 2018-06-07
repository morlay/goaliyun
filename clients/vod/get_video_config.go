package vod

import (
	"github.com/morlay/goaliyun"
)

type GetVideoConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetVideoConfigRequest) Invoke(client goaliyun.Client) (*GetVideoConfigResponse, error) {
	resp := &GetVideoConfigResponse{}
	err := client.Request("vod", "GetVideoConfig", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetVideoConfigResponse struct {
	RequestId      goaliyun.String
	DownloadSwitch goaliyun.String
}
