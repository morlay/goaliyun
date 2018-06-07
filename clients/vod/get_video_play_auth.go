package vod

import (
	"github.com/morlay/goaliyun"
)

type GetVideoPlayAuthRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ReAuthInfo           string `position:"Query" name:"ReAuthInfo"`
	AuthInfoTimeout      int64  `position:"Query" name:"AuthInfoTimeout"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetVideoPlayAuthRequest) Invoke(client goaliyun.Client) (*GetVideoPlayAuthResponse, error) {
	resp := &GetVideoPlayAuthResponse{}
	err := client.Request("vod", "GetVideoPlayAuth", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetVideoPlayAuthResponse struct {
	RequestId goaliyun.String
	PlayAuth  goaliyun.String
	VideoMeta GetVideoPlayAuthVideoMeta
}

type GetVideoPlayAuthVideoMeta struct {
	CoverURL goaliyun.String
	Duration goaliyun.Float
	Status   goaliyun.String
	Title    goaliyun.String
	VideoId  goaliyun.String
}
