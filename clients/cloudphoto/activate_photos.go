package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ActivatePhotosRequest struct {
	LibraryId string                     `position:"Query" name:"LibraryId"`
	PhotoIds  *ActivatePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                     `position:"Query" name:"StoreName"`
	RegionId  string                     `position:"Query" name:"RegionId"`
}

func (req *ActivatePhotosRequest) Invoke(client goaliyun.Client) (*ActivatePhotosResponse, error) {
	resp := &ActivatePhotosResponse{}
	err := client.Request("cloudphoto", "ActivatePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActivatePhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   ActivatePhotosResultList
}

type ActivatePhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type ActivatePhotosPhotoIdList []int64

func (list *ActivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type ActivatePhotosResultList []ActivatePhotosResult

func (list *ActivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ActivatePhotosResult)
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
