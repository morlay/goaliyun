package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPhotosRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListPhotosRequest) Invoke(client goaliyun.Client) (*ListPhotosResponse, error) {
	resp := &ListPhotosResponse{}
	err := client.Request("cloudphoto", "ListPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Photos     ListPhotosPhotoList
}

type ListPhotosPhoto struct {
	Id              goaliyun.Integer
	IdStr           goaliyun.String
	Title           goaliyun.String
	FileId          goaliyun.String
	Location        goaliyun.String
	State           goaliyun.String
	Md5             goaliyun.String
	IsVideo         bool
	Remark          goaliyun.String
	Size            goaliyun.Integer
	Width           goaliyun.Integer
	Height          goaliyun.Integer
	Ctime           goaliyun.Integer
	Mtime           goaliyun.Integer
	TakenAt         goaliyun.Integer
	InactiveTime    goaliyun.Integer
	ShareExpireTime goaliyun.Integer
}

type ListPhotosPhotoList []ListPhotosPhoto

func (list *ListPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotosPhoto)
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
