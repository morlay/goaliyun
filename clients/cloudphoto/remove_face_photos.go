package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveFacePhotosRequest struct {
	LibraryId string                       `position:"Query" name:"LibraryId"`
	PhotoIds  *RemoveFacePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                       `position:"Query" name:"StoreName"`
	FaceId    int64                        `position:"Query" name:"FaceId"`
	RegionId  string                       `position:"Query" name:"RegionId"`
}

func (req *RemoveFacePhotosRequest) Invoke(client goaliyun.Client) (*RemoveFacePhotosResponse, error) {
	resp := &RemoveFacePhotosResponse{}
	err := client.Request("cloudphoto", "RemoveFacePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveFacePhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   RemoveFacePhotosResultList
}

type RemoveFacePhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type RemoveFacePhotosPhotoIdList []int64

func (list *RemoveFacePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type RemoveFacePhotosResultList []RemoveFacePhotosResult

func (list *RemoveFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveFacePhotosResult)
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
