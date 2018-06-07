package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type GetVideoCoverRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	ZoomType  string `position:"Query" name:"ZoomType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetVideoCoverRequest) Invoke(client goaliyun.Client) (*GetVideoCoverResponse, error) {
	resp := &GetVideoCoverResponse{}
	err := client.Request("cloudphoto", "GetVideoCover", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetVideoCoverResponse struct {
	Code          goaliyun.String
	Message       goaliyun.String
	VideoCoverUrl goaliyun.String
	RequestId     goaliyun.String
	Action        goaliyun.String
}
