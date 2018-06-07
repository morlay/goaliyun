package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type GetEventRequest struct {
	EventId   int64  `position:"Query" name:"EventId"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetEventRequest) Invoke(client goaliyun.Client) (*GetEventResponse, error) {
	resp := &GetEventResponse{}
	err := client.Request("cloudphoto", "GetEvent", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetEventResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Event     GetEventEvent
}

type GetEventEvent struct {
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
