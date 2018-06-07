package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ReactivatePhotosRequest struct {
	LibraryId string                       `position:"Query" name:"LibraryId"`
	PhotoIds  *ReactivatePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                       `position:"Query" name:"StoreName"`
	RegionId  string                       `position:"Query" name:"RegionId"`
}

func (req *ReactivatePhotosRequest) Invoke(client goaliyun.Client) (*ReactivatePhotosResponse, error) {
	resp := &ReactivatePhotosResponse{}
	err := client.Request("cloudphoto", "ReactivatePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReactivatePhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   ReactivatePhotosResultList
}

type ReactivatePhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type ReactivatePhotosPhotoIdList []int64

func (list *ReactivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type ReactivatePhotosResultList []ReactivatePhotosResult

func (list *ReactivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ReactivatePhotosResult)
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
