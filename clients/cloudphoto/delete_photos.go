package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeletePhotosRequest struct {
	LibraryId string                   `position:"Query" name:"LibraryId"`
	StoreName string                   `position:"Query" name:"StoreName"`
	PhotoIds  *DeletePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	RegionId  string                   `position:"Query" name:"RegionId"`
}

func (req *DeletePhotosRequest) Invoke(client goaliyun.Client) (*DeletePhotosResponse, error) {
	resp := &DeletePhotosResponse{}
	err := client.Request("cloudphoto", "DeletePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   DeletePhotosResultList
}

type DeletePhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type DeletePhotosPhotoIdList []int64

func (list *DeletePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type DeletePhotosResultList []DeletePhotosResult

func (list *DeletePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeletePhotosResult)
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
