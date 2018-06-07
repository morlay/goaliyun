package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAlbumsRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListAlbumsRequest) Invoke(client goaliyun.Client) (*ListAlbumsResponse, error) {
	resp := &ListAlbumsResponse{}
	err := client.Request("cloudphoto", "ListAlbums", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAlbumsResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Albums     ListAlbumsAlbumList
}

type ListAlbumsAlbum struct {
	Id          goaliyun.Integer
	IdStr       goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	Remark      goaliyun.String
	PhotosCount goaliyun.Integer
	Ctime       goaliyun.Integer
	Mtime       goaliyun.Integer
	Cover       ListAlbumsCover
}

type ListAlbumsCover struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Title   goaliyun.String
	FileId  goaliyun.String
	State   goaliyun.String
	Md5     goaliyun.String
	IsVideo bool
	Remark  goaliyun.String
	Width   goaliyun.Integer
	Height  goaliyun.Integer
	Ctime   goaliyun.Integer
	Mtime   goaliyun.Integer
}

type ListAlbumsAlbumList []ListAlbumsAlbum

func (list *ListAlbumsAlbumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlbumsAlbum)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
