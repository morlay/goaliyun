package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTimeLinePhotosRequest struct {
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	EndTime   int64  `position:"Query" name:"EndTime"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int64  `position:"Query" name:"Page"`
	StartTime int64  `position:"Query" name:"StartTime"`
	FilterBy  string `position:"Query" name:"FilterBy"`
	Direction string `position:"Query" name:"Direction"`
	Order     string `position:"Query" name:"Order"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListTimeLinePhotosRequest) Invoke(client goaliyun.Client) (*ListTimeLinePhotosResponse, error) {
	resp := &ListTimeLinePhotosResponse{}
	err := client.Request("cloudphoto", "ListTimeLinePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTimeLinePhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Photos     ListTimeLinePhotosPhotoList
}

type ListTimeLinePhotosPhoto struct {
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
	ShareExpireTime goaliyun.Integer
	Like            goaliyun.Integer
}

type ListTimeLinePhotosPhotoList []ListTimeLinePhotosPhoto

func (list *ListTimeLinePhotosPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTimeLinePhotosPhoto)
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
