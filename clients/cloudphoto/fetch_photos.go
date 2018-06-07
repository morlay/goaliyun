package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FetchPhotosRequest struct {
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	OrderBy   string `position:"Query" name:"OrderBy"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Page      int64  `position:"Query" name:"Page"`
	Order     string `position:"Query" name:"Order"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *FetchPhotosRequest) Invoke(client goaliyun.Client) (*FetchPhotosResponse, error) {
	resp := &FetchPhotosResponse{}
	err := client.Request("cloudphoto", "FetchPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FetchPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Photos     FetchPhotosPhotoList
}

type FetchPhotosPhoto struct {
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

type FetchPhotosPhotoList []FetchPhotosPhoto

func (list *FetchPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchPhotosPhoto)
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
