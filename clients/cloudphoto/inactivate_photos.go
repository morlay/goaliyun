package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InactivatePhotosRequest struct {
	LibraryId    string                       `position:"Query" name:"LibraryId"`
	PhotoIds     *InactivatePhotosPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName    string                       `position:"Query" name:"StoreName"`
	InactiveTime int64                        `position:"Query" name:"InactiveTime"`
	RegionId     string                       `position:"Query" name:"RegionId"`
}

func (req *InactivatePhotosRequest) Invoke(client goaliyun.Client) (*InactivatePhotosResponse, error) {
	resp := &InactivatePhotosResponse{}
	err := client.Request("cloudphoto", "InactivatePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InactivatePhotosResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Results   InactivatePhotosResultList
}

type InactivatePhotosResult struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Code    goaliyun.String
	Message goaliyun.String
}

type InactivatePhotosPhotoIdList []int64

func (list *InactivatePhotosPhotoIdList) UnmarshalJSON(data []byte) error {
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

type InactivatePhotosResultList []InactivatePhotosResult

func (list *InactivatePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InactivatePhotosResult)
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
