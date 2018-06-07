package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MoveAlbumPhotosRequest struct {
	SourceAlbumId int64                       `position:"Query" name:"SourceAlbumId"`
	TargetAlbumId int64                       `position:"Query" name:"TargetAlbumId"`
	LibraryId     string                      `position:"Query" name:"LibraryId"`
	PhotoIds      *MoveAlbumPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName     string                      `position:"Query" name:"StoreName"`
	RegionId      string                      `position:"Query" name:"RegionId"`
}

func (req *MoveAlbumPhotosRequest) Invoke(client goaliyun.Client) (*MoveAlbumPhotosResponse, error) {
	resp := &MoveAlbumPhotosResponse{}
	err := client.Request("cloudphoto", "MoveAlbumPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MoveAlbumPhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   MoveAlbumPhotosResultList
}

type MoveAlbumPhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type MoveAlbumPhotosPhotoIdList []int64

func (list *MoveAlbumPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type MoveAlbumPhotosResultList []MoveAlbumPhotosResult

func (list *MoveAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MoveAlbumPhotosResult)
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
