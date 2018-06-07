package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type CreateAlbumRequest struct {
	AlbumName string `position:"Query" name:"AlbumName"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Remark    string `position:"Query" name:"Remark"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *CreateAlbumRequest) Invoke(client goaliyun.Client) (*CreateAlbumResponse, error) {
	resp := &CreateAlbumResponse{}
	err := client.Request("cloudphoto", "CreateAlbum", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAlbumResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Album     CreateAlbumAlbum
}

type CreateAlbumAlbum struct {
	Id          goaliyun.Integer
	IdStr       goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Remark      goaliyun.String
	PhotosCount goaliyun.Integer
	Ctime       goaliyun.Integer
	Mtime       goaliyun.Integer
	Cover       CreateAlbumCover
}

type CreateAlbumCover struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Title   goaliyun.String
	FileId  goaliyun.String
	State   goaliyun.String
	Md5     goaliyun.String
	IsVideo bool
	Width   goaliyun.Integer
	Height  goaliyun.Integer
	Ctime   goaliyun.Integer
	Mtime   goaliyun.Integer
	Remark  goaliyun.String
}
