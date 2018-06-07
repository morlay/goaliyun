package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetPhotosRequest struct {
	LibraryId string                `position:"Query" name:"LibraryId"`
	PhotoIds  *GetPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                `position:"Query" name:"StoreName"`
	RegionId  string                `position:"Query" name:"RegionId"`
}

func (req *GetPhotosRequest) Invoke(client goaliyun.Client) (*GetPhotosResponse, error) {
	resp := &GetPhotosResponse{}
	err := client.Request("cloudphoto", "GetPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Photos    GetPhotosPhotoList
}

type GetPhotosPhoto struct {
	Id              goaliyun.Integer
	IdStr           goaliyun.String
	Title           goaliyun.String
	FileId          goaliyun.String
	Location        goaliyun.String
	State           goaliyun.String
	Md5             goaliyun.String
	IsVideo         bool
	Remark          goaliyun.String
	Width           goaliyun.Integer
	Height          goaliyun.Integer
	Size            goaliyun.Integer
	Ctime           goaliyun.Integer
	Mtime           goaliyun.Integer
	TakenAt         goaliyun.Integer
	InactiveTime    goaliyun.Integer
	ShareExpireTime goaliyun.Integer
	Like            goaliyun.Integer
}

type GetPhotosPhotoIdList []int64

func (list *GetPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetPhotosPhotoList []GetPhotosPhoto

func (list *GetPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotosPhoto)
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
