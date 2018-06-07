package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTimeLinesRequest struct {
	Cursor        int64  `position:"Query" name:"Cursor"`
	PhotoSize     int64  `position:"Query" name:"PhotoSize"`
	TimeLineCount int64  `position:"Query" name:"TimeLineCount"`
	LibraryId     string `position:"Query" name:"LibraryId"`
	StoreName     string `position:"Query" name:"StoreName"`
	TimeLineUnit  string `position:"Query" name:"TimeLineUnit"`
	FilterBy      string `position:"Query" name:"FilterBy"`
	Direction     string `position:"Query" name:"Direction"`
	Order         string `position:"Query" name:"Order"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ListTimeLinesRequest) Invoke(client goaliyun.Client) (*ListTimeLinesResponse, error) {
	resp := &ListTimeLinesResponse{}
	err := client.Request("cloudphoto", "ListTimeLines", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTimeLinesResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	TimeLines  ListTimeLinesTimeLineList
}

type ListTimeLinesTimeLine struct {
	StartTime   goaliyun.Integer
	EndTime     goaliyun.Integer
	TotalCount  goaliyun.Integer
	PhotosCount goaliyun.Integer
	Photos      ListTimeLinesPhotoList
}

type ListTimeLinesPhoto struct {
	Id              goaliyun.Integer
	IdStr           goaliyun.String
	Title           goaliyun.String
	Location        goaliyun.String
	FileId          goaliyun.String
	State           goaliyun.String
	Md5             goaliyun.String
	IsVideo         bool
	Remark          goaliyun.String
	Size            goaliyun.Integer
	Width           goaliyun.Integer
	Height          goaliyun.Integer
	Ctime           goaliyun.Integer
	Mtime           goaliyun.Integer
	TakenAt         goaliyun.Integer
	ShareExpireTime goaliyun.Integer
	Like            goaliyun.Integer
}

type ListTimeLinesTimeLineList []ListTimeLinesTimeLine

func (list *ListTimeLinesTimeLineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTimeLinesTimeLine)
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

type ListTimeLinesPhotoList []ListTimeLinesPhoto

func (list *ListTimeLinesPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTimeLinesPhoto)
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
