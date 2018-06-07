package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type GetDownloadUrlRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetDownloadUrlRequest) Invoke(client goaliyun.Client) (*GetDownloadUrlResponse, error) {
	resp := &GetDownloadUrlResponse{}
	err := client.Request("cloudphoto", "GetDownloadUrl", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetDownloadUrlResponse struct {
	Code        goaliyun.String
	Message     goaliyun.String
	DownloadUrl goaliyun.String
	RequestId   goaliyun.String
	Action      goaliyun.String
}
