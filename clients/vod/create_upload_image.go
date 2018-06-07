package vod

import (
	"github.com/morlay/goaliyun"
)

type CreateUploadImageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageType            string `position:"Query" name:"ImageType"`
	OriginalFileName     string `position:"Query" name:"OriginalFileName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ImageExt             string `position:"Query" name:"ImageExt"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateUploadImageRequest) Invoke(client goaliyun.Client) (*CreateUploadImageResponse, error) {
	resp := &CreateUploadImageResponse{}
	err := client.Request("vod", "CreateUploadImage", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateUploadImageResponse struct {
	RequestId     goaliyun.String
	ImageId       goaliyun.String
	ImageURL      goaliyun.String
	UploadAddress goaliyun.String
	UploadAuth    goaliyun.String
}
