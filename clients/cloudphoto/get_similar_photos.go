package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetSimilarPhotosRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetSimilarPhotosRequest) Invoke(client goaliyun.Client) (*GetSimilarPhotosResponse, error) {
	resp := &GetSimilarPhotosResponse{}
	err := client.Request("cloudphoto", "GetSimilarPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetSimilarPhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Photos    GetSimilarPhotosPhotoList
}

type GetSimilarPhotosPhoto struct {
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
	Like            goaliyun.Integer
}

type GetSimilarPhotosPhotoList []GetSimilarPhotosPhoto

func (list *GetSimilarPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSimilarPhotosPhoto)
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
