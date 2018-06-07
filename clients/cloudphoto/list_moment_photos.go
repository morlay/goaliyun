package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMomentPhotosRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	MomentId  int64  `position:"Query" name:"MomentId"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListMomentPhotosRequest) Invoke(client goaliyun.Client) (*ListMomentPhotosResponse, error) {
	resp := &ListMomentPhotosResponse{}
	err := client.Request("cloudphoto", "ListMomentPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMomentPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Results    ListMomentPhotosResultList
}

type ListMomentPhotosResult struct {
	PhotoId    goaliyun.Integer
	PhotoIdStr goaliyun.String
	State      goaliyun.String
}

type ListMomentPhotosResultList []ListMomentPhotosResult

func (list *ListMomentPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMomentPhotosResult)
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
