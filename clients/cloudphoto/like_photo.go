package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type LikePhotoRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *LikePhotoRequest) Invoke(client goaliyun.Client) (*LikePhotoResponse, error) {
	resp := &LikePhotoResponse{}
	err := client.Request("cloudphoto", "LikePhoto", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type LikePhotoResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}
