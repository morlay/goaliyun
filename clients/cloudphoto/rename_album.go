package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type RenameAlbumRequest struct {
	AlbumName string `position:"Query" name:"AlbumName"`
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *RenameAlbumRequest) Invoke(client goaliyun.Client) (*RenameAlbumResponse, error) {
	resp := &RenameAlbumResponse{}
	err := client.Request("cloudphoto", "RenameAlbum", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenameAlbumResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}
