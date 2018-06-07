package vod

import (
	"github.com/morlay/goaliyun"
)

type GetMezzanineInfoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	PreviewSegment       string `position:"Query" name:"PreviewSegment"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetMezzanineInfoRequest) Invoke(client goaliyun.Client) (*GetMezzanineInfoResponse, error) {
	resp := &GetMezzanineInfoResponse{}
	err := client.Request("vod", "GetMezzanineInfo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetMezzanineInfoResponse struct {
	RequestId goaliyun.String
	Mezzanine GetMezzanineInfoMezzanine
}

type GetMezzanineInfoMezzanine struct {
	VideoId          goaliyun.String
	Bitrate          goaliyun.String
	CreationTime     goaliyun.String
	Duration         goaliyun.String
	Fps              goaliyun.String
	Height           goaliyun.Integer
	Width            goaliyun.Integer
	Size             goaliyun.Integer
	Status           goaliyun.String
	FileURL          goaliyun.String
	FileName         goaliyun.String
	CRC64            goaliyun.String
	PreprocessStatus goaliyun.String
}
