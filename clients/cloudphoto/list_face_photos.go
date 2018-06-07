package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFacePhotosRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListFacePhotosRequest) Invoke(client goaliyun.Client) (*ListFacePhotosResponse, error) {
	resp := &ListFacePhotosResponse{}
	err := client.Request("cloudphoto", "ListFacePhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFacePhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Results    ListFacePhotosResultList
}

type ListFacePhotosResult struct {
	PhotoId    goaliyun.Integer
	PhotoIdStr goaliyun.String
	Mtime      goaliyun.Integer
	State      goaliyun.String
}

type ListFacePhotosResultList []ListFacePhotosResult

func (list *ListFacePhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFacePhotosResult)
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
