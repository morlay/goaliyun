package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type SetAlbumCoverRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *SetAlbumCoverRequest) Invoke(client goaliyun.Client) (*SetAlbumCoverResponse, error) {
	resp := &SetAlbumCoverResponse{}
	err := client.Request("cloudphoto", "SetAlbumCover", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetAlbumCoverResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}
