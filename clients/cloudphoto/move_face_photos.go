package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type MoveFacePhotosRequest struct {
	LibraryId    string                     `position:"Query" name:"LibraryId"`
	TargetFaceId int64                      `position:"Query" name:"TargetFaceId"`
	PhotoIds     *MoveFacePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName    string                     `position:"Query" name:"StoreName"`
	SourceFaceId int64                      `position:"Query" name:"SourceFaceId"`
	RegionId     string                     `position:"Query" name:"RegionId"`
}

func (req *MoveFacePhotosRequest) Invoke(client goaliyun.Client) (*MoveFacePhotosResponse, error) {
	resp := &MoveFacePhotosResponse{}
	err := client.Request("cloudphoto", "MoveFacePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MoveFacePhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   MoveFacePhotosResultList
}

type MoveFacePhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type MoveFacePhotosPhotoIdList []int64

func (list *MoveFacePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type MoveFacePhotosResultList []MoveFacePhotosResult

func (list *MoveFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]MoveFacePhotosResult)
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
