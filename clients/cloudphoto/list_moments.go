package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMomentsRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListMomentsRequest) Invoke(client goaliyun.Client) (*ListMomentsResponse, error) {
	resp := &ListMomentsResponse{}
	err := client.Request("cloudphoto", "ListMoments", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMomentsResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Moments    ListMomentsMomentList
}

type ListMomentsMoment struct {
	Id           goaliyun.Integer
	IdStr        goaliyun.String
	LocationName goaliyun.String
	PhotosCount  goaliyun.Integer
	State        goaliyun.String
	TakenAt      goaliyun.Integer
	Ctime        goaliyun.Integer
	Mtime        goaliyun.Integer
}

type ListMomentsMomentList []ListMomentsMoment

func (list *ListMomentsMomentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMomentsMoment)
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
