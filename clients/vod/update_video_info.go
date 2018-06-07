package vod

import (
	"github.com/morlay/goaliyun"
)

type UpdateVideoInfoRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	Description          string `position:"Query" name:"Description"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateVideoInfoRequest) Invoke(client goaliyun.Client) (*UpdateVideoInfoResponse, error) {
	resp := &UpdateVideoInfoResponse{}
	err := client.Request("vod", "UpdateVideoInfo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateVideoInfoResponse struct {
	RequestId goaliyun.String
}
