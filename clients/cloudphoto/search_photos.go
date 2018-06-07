package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchPhotosRequest struct {
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int64  `position:"Query" name:"Page"`
	Keyword   string `position:"Query" name:"Keyword"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *SearchPhotosRequest) Invoke(client goaliyun.Client) (*SearchPhotosResponse, error) {
	resp := &SearchPhotosResponse{}
	err := client.Request("cloudphoto", "SearchPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Photos     SearchPhotosPhotoList
}

type SearchPhotosPhoto struct {
	Id              goaliyun.Integer
	IdStr           goaliyun.String
	Title           goaliyun.String
	FileId          goaliyun.String
	Location        goaliyun.String
	State           goaliyun.String
	Md5             goaliyun.String
	IsVideo         bool
	Size            goaliyun.Integer
	Width           goaliyun.Integer
	Height          goaliyun.Integer
	Ctime           goaliyun.Integer
	Mtime           goaliyun.Integer
	TakenAt         goaliyun.Integer
	ShareExpireTime goaliyun.Integer
}

type SearchPhotosPhotoList []SearchPhotosPhoto

func (list *SearchPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchPhotosPhoto)
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
