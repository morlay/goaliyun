package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddAlbumPhotosRequest struct {
	LibraryId string                     `position:"Query" name:"LibraryId"`
	AlbumId   int64                      `position:"Query" name:"AlbumId"`
	PhotoIds  *AddAlbumPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                     `position:"Query" name:"StoreName"`
	RegionId  string                     `position:"Query" name:"RegionId"`
}

func (req *AddAlbumPhotosRequest) Invoke(client goaliyun.Client) (*AddAlbumPhotosResponse, error) {
	resp := &AddAlbumPhotosResponse{}
	err := client.Request("cloudphoto", "AddAlbumPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddAlbumPhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   AddAlbumPhotosResultList
}

type AddAlbumPhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type AddAlbumPhotosPhotoIdList []int64

func (list *AddAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type AddAlbumPhotosResultList []AddAlbumPhotosResult

func (list *AddAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddAlbumPhotosResult)
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
