package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type CreateEventRequest struct {
	BannerPhotoId    string `position:"Query" name:"BannerPhotoId"`
	WatermarkPhotoId string `position:"Query" name:"WatermarkPhotoId"`
	Identity         string `position:"Query" name:"Identity"`
	SplashPhotoId    string `position:"Query" name:"SplashPhotoId"`
	LibraryId        string `position:"Query" name:"LibraryId"`
	WeixinTitle      string `position:"Query" name:"WeixinTitle"`
	StoreName        string `position:"Query" name:"StoreName"`
	Remark           string `position:"Query" name:"Remark"`
	Title            string `position:"Query" name:"Title"`
	EndAt            int64  `position:"Query" name:"EndAt"`
	StartAt          int64  `position:"Query" name:"StartAt"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *CreateEventRequest) Invoke(client goaliyun.Client) (*CreateEventResponse, error) {
	resp := &CreateEventResponse{}
	err := client.Request("cloudphoto", "CreateEvent", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateEventResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Event     CreateEventEvent
}

type CreateEventEvent struct {
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
