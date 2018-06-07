package vod

import (
	"github.com/morlay/goaliyun"
)

type ProduceEditingProjectVideoRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MediaMetadata        string `position:"Query" name:"MediaMetadata"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Timeline             string `position:"Query" name:"Timeline"`
	Description          string `position:"Query" name:"Description"`
	ProduceConfig        string `position:"Query" name:"ProduceConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	ProjectId            string `position:"Query" name:"ProjectId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ProduceEditingProjectVideoRequest) Invoke(client goaliyun.Client) (*ProduceEditingProjectVideoResponse, error) {
	resp := &ProduceEditingProjectVideoResponse{}
	err := client.Request("vod", "ProduceEditingProjectVideo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ProduceEditingProjectVideoResponse struct {
	RequestId goaliyun.String
	MediaId   goaliyun.String
	ProjectId goaliyun.String
}
