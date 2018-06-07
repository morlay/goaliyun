package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type EditPhotosRequest struct {
	TakenAt         int64                  `position:"Query" name:"TakenAt"`
	LibraryId       string                 `position:"Query" name:"LibraryId"`
	ShareExpireTime int64                  `position:"Query" name:"ShareExpireTime"`
	PhotoIds        *EditPhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName       string                 `position:"Query" name:"StoreName"`
	Remark          string                 `position:"Query" name:"Remark"`
	Title           string                 `position:"Query" name:"Title"`
	RegionId        string                 `position:"Query" name:"RegionId"`
}

func (req *EditPhotosRequest) Invoke(client goaliyun.Client) (*EditPhotosResponse, error) {
	resp := &EditPhotosResponse{}
	err := client.Request("cloudphoto", "EditPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EditPhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   EditPhotosResultList
}

type EditPhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type EditPhotosPhotoIdList []int64

func (list *EditPhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type EditPhotosResultList []EditPhotosResult

func (list *EditPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]EditPhotosResult)
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
