package vod

import (
	"github.com/morlay/goaliyun"
)

type GetVideoPlayInfoRequest struct {
	SignVersion          string `position:"Query" name:"SignVersion"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientVersion        string `position:"Query" name:"ClientVersion"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Channel              string `position:"Query" name:"Channel"`
	PlaySign             string `position:"Query" name:"PlaySign"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ClientTS             int64  `position:"Query" name:"ClientTS"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetVideoPlayInfoRequest) Invoke(client goaliyun.Client) (*GetVideoPlayInfoResponse, error) {
	resp := &GetVideoPlayInfoResponse{}
	err := client.Request("vod", "GetVideoPlayInfo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetVideoPlayInfoResponse struct {
	RequestId goaliyun.String
	PlayInfo  GetVideoPlayInfoPlayInfo
	VideoInfo GetVideoPlayInfoVideoInfo
}

type GetVideoPlayInfoPlayInfo struct {
	AccessKeyId     goaliyun.String
	AccessKeySecret goaliyun.String
	AuthInfo        goaliyun.String
	SecurityToken   goaliyun.String
	Region          goaliyun.String
	PlayDomain      goaliyun.String
}

type GetVideoPlayInfoVideoInfo struct {
	CoverURL   goaliyun.String
	CustomerId goaliyun.Integer
	Duration   goaliyun.Float
	Status     goaliyun.String
	Title      goaliyun.String
	VideoId    goaliyun.String
}
