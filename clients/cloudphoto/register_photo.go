package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type RegisterPhotoRequest struct {
	LibraryId  string  `position:"Query" name:"LibraryId"`
	Latitude   float64 `position:"Query" name:"Latitude"`
	PhotoTitle string  `position:"Query" name:"PhotoTitle"`
	StoreName  string  `position:"Query" name:"StoreName"`
	IsVideo    string  `position:"Query" name:"IsVideo"`
	Remark     string  `position:"Query" name:"Remark"`
	Size       int64   `position:"Query" name:"Size"`
	TakenAt    int64   `position:"Query" name:"TakenAt"`
	Width      int64   `position:"Query" name:"Width"`
	Location   string  `position:"Query" name:"Location"`
	Longitude  float64 `position:"Query" name:"Longitude"`
	Height     int64   `position:"Query" name:"Height"`
	Md5        string  `position:"Query" name:"Md.5"`
	RegionId   string  `position:"Query" name:"RegionId"`
}

func (req *RegisterPhotoRequest) Invoke(client goaliyun.Client) (*RegisterPhotoResponse, error) {
	resp := &RegisterPhotoResponse{}
	err := client.Request("cloudphoto", "RegisterPhoto", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RegisterPhotoResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Photo     RegisterPhotoPhoto
}

type RegisterPhotoPhoto struct {
	Id              goaliyun.Integer
	IdStr           goaliyun.String
	Title           goaliyun.String
	Location        goaliyun.String
	FileId          goaliyun.String
	State           goaliyun.String
	Md5             goaliyun.String
	IsVideo         bool
	Size            goaliyun.Integer
	Remark          goaliyun.String
	Width           goaliyun.Integer
	Height          goaliyun.Integer
	Ctime           goaliyun.Integer
	Mtime           goaliyun.Integer
	TakenAt         goaliyun.Integer
	ShareExpireTime goaliyun.Integer
}
