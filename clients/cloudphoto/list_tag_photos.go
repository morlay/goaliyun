package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagPhotosRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	TagId     int64  `position:"Query" name:"TagId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListTagPhotosRequest) Invoke(client goaliyun.Client) (*ListTagPhotosResponse, error) {
	resp := &ListTagPhotosResponse{}
	err := client.Request("cloudphoto", "ListTagPhotos", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagPhotosResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Results    ListTagPhotosResultList
}

type ListTagPhotosResult struct {
	PhotoId    goaliyun.Integer
	PhotoIdStr goaliyun.String
	State      goaliyun.String
}

type ListTagPhotosResultList []ListTagPhotosResult

func (list *ListTagPhotosResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagPhotosResult)
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
