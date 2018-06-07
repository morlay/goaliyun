package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetAlbumsByNamesRequest struct {
	LibraryId string                    `position:"Query" name:"LibraryId"`
	Names     *GetAlbumsByNamesNameList `position:"Query" type:"Repeated" name:"Name"`
	StoreName string                    `position:"Query" name:"StoreName"`
	RegionId  string                    `position:"Query" name:"RegionId"`
}

func (req *GetAlbumsByNamesRequest) Invoke(client goaliyun.Client) (*GetAlbumsByNamesResponse, error) {
	resp := &GetAlbumsByNamesResponse{}
	err := client.Request("cloudphoto", "GetAlbumsByNames", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAlbumsByNamesResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Albums    GetAlbumsByNamesAlbumList
}

type GetAlbumsByNamesAlbum struct {
	Id          goaliyun.Integer
	IdStr       goaliyun.String
	Name        goaliyun.String
	State       goaliyun.String
	PhotosCount goaliyun.Integer
	Ctime       goaliyun.Integer
	Mtime       goaliyun.Integer
	Cover       GetAlbumsByNamesCover
}

type GetAlbumsByNamesCover struct {
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

type GetAlbumsByNamesNameList []string

func (list *GetAlbumsByNamesNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type GetAlbumsByNamesAlbumList []GetAlbumsByNamesAlbum

func (list *GetAlbumsByNamesAlbumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAlbumsByNamesAlbum)
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
