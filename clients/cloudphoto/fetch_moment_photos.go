package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FetchMomentPhotosRequest struct {
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	OrderBy   string `position:"Query" name:"OrderBy"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int64  `position:"Query" name:"Page"`
	MomentId  int64  `position:"Query" name:"MomentId"`
	Order     string `position:"Query" name:"Order"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *FetchMomentPhotosRequest) Invoke(client goaliyun.Client) (*FetchMomentPhotosResponse, error) {
	resp := &FetchMomentPhotosResponse{}
	err := client.Request("cloudphoto", "FetchMomentPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FetchMomentPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Photos     FetchMomentPhotosPhotoList
}

type FetchMomentPhotosPhoto struct {
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

type FetchMomentPhotosPhotoList []FetchMomentPhotosPhoto

func (list *FetchMomentPhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchMomentPhotosPhoto)
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
