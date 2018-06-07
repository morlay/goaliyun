package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAlbumPhotosRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListAlbumPhotosRequest) Invoke(client goaliyun.Client) (*ListAlbumPhotosResponse, error) {
	resp := &ListAlbumPhotosResponse{}
	err := client.Request("cloudphoto", "ListAlbumPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAlbumPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Results    ListAlbumPhotosResultList
}

type ListAlbumPhotosResult struct {
	PhotoId    goaliyun.Integer
	PhotoIdStr goaliyun.String
	Mtime      goaliyun.Integer
	State      goaliyun.String
}

type ListAlbumPhotosResultList []ListAlbumPhotosResult

func (list *ListAlbumPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlbumPhotosResult)
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
