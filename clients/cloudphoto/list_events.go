package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListEventsRequest struct {
	Cursor    string `position:"Query" name:"Cursor"`
	Size      int64  `position:"Query" name:"Size"`
	StoreName string `position:"Query" name:"StoreName"`
	State     string `position:"Query" name:"State"`
	Direction string `position:"Query" name:"Direction"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListEventsRequest) Invoke(client goaliyun.Client) (*ListEventsResponse, error) {
	resp := &ListEventsResponse{}
	err := client.Request("cloudphoto", "ListEvents", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListEventsResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Events     ListEventsEventList
}

type ListEventsEvent struct {
	Id               goaliyun.Integer
	IdStr            goaliyun.String
	Title            goaliyun.String
	BannerPhotoId    goaliyun.String
	Identity         goaliyun.String
	SplashPhotoId    goaliyun.String
	State            goaliyun.String
	WeixinTitle      goaliyun.String
	WatermarkPhotoId goaliyun.String
	StartAt          goaliyun.Integer
	EndAt            goaliyun.Integer
	Ctime            goaliyun.Integer
	Mtime            goaliyun.Integer
	ViewsCount       goaliyun.Integer
	LibraryId        goaliyun.String
	IdStr1           goaliyun.String
}

type ListEventsEventList []ListEventsEvent

func (list *ListEventsEventList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEventsEvent)
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
