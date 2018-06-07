package vod

import (
	"github.com/morlay/goaliyun"
)

type CreateUploadVideoRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TranscodeMode        string `position:"Query" name:"TranscodeMode"`
	IP                   string `position:"Query" name:"IP"`
	Description          string `position:"Query" name:"Description"`
	FileSize             int64  `position:"Query" name:"FileSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
	StorageLocation      string `position:"Query" name:"StorageLocation"`
	CoverURL             string `position:"Query" name:"CoverURL"`
	UserData             string `position:"Query" name:"UserData"`
	FileName             string `position:"Query" name:"FileName"`
	TemplateGroupId      string `position:"Query" name:"TemplateGroupId"`
	CateId               int64  `position:"Query" name:"CateId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateUploadVideoRequest) Invoke(client goaliyun.Client) (*CreateUploadVideoResponse, error) {
	resp := &CreateUploadVideoResponse{}
	err := client.Request("vod", "CreateUploadVideo", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateUploadVideoResponse struct {
	RequestId     goaliyun.String
	VideoId       goaliyun.String
	UploadAddress goaliyun.String
	UploadAuth    goaliyun.String
}
