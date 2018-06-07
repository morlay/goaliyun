package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveAlbumPhotosRequest struct {
	LibraryId string                        `position:"Query" name:"LibraryId"`
	AlbumId   int64                         `position:"Query" name:"AlbumId"`
	PhotoIds  *RemoveAlbumPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                        `position:"Query" name:"StoreName"`
	RegionId  string                        `position:"Query" name:"RegionId"`
}

func (req *RemoveAlbumPhotosRequest) Invoke(client goaliyun.Client) (*RemoveAlbumPhotosResponse, error) {
	resp := &RemoveAlbumPhotosResponse{}
	err := client.Request("cloudphoto", "RemoveAlbumPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveAlbumPhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   RemoveAlbumPhotosResultList
}

type RemoveAlbumPhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type RemoveAlbumPhotosPhotoIdList []int64

func (list *RemoveAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type RemoveAlbumPhotosResultList []RemoveAlbumPhotosResult

func (list *RemoveAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveAlbumPhotosResult)
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
