package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FetchAlbumTagPhotosRequest struct {
	Size      int64  `position:"Query" name:"Size"`
	TagId     int64  `position:"Query" name:"TagId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	AlbumId   int64  `position:"Query" name:"AlbumId"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int64  `position:"Query" name:"Page"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *FetchAlbumTagPhotosRequest) Invoke(client goaliyun.Client) (*FetchAlbumTagPhotosResponse, error) {
	resp := &FetchAlbumTagPhotosResponse{}
	err := client.Request("cloudphoto", "FetchAlbumTagPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FetchAlbumTagPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Results    FetchAlbumTagPhotosResultList
}

type FetchAlbumTagPhotosResult struct {
	PhotoId    goaliyun.Integer
	PhotoIdStr goaliyun.String
	Mtime      goaliyun.Integer
	State      goaliyun.String
}

type FetchAlbumTagPhotosResultList []FetchAlbumTagPhotosResult

func (list *FetchAlbumTagPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchAlbumTagPhotosResult)
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
