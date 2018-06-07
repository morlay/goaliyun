package vod

import (
	"github.com/morlay/goaliyun"
)

type GetImageInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetImageInfoRequest) Invoke(client goaliyun.Client) (*GetImageInfoResponse, error) {
	resp := &GetImageInfoResponse{}
	err := client.Request("vod", "GetImageInfo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetImageInfoResponse struct {
	RequestId goaliyun.String
	ImageInfo GetImageInfoImageInfo
}

type GetImageInfoImageInfo struct {
	ImageId      goaliyun.String
	Title        goaliyun.String
	CreationTime goaliyun.String
	ImageType    goaliyun.String
	Tags         goaliyun.String
	URL          goaliyun.String
	Mezzanine    GetImageInfoMezzanine
}

type GetImageInfoMezzanine struct {
	OriginalFileName goaliyun.String
}
