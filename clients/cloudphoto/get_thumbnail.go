package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type GetThumbnailRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	ZoomType  string `position:"Query" name:"ZoomType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetThumbnailRequest) Invoke(client goaliyun.Client) (*GetThumbnailResponse, error) {
	resp := &GetThumbnailResponse{}
	err := client.Request("cloudphoto", "GetThumbnail", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetThumbnailResponse struct {
	Code         goaliyun.String
	Message      goaliyun.String
	ThumbnailUrl goaliyun.String
	RequestId    goaliyun.String
	Action       goaliyun.String
}
