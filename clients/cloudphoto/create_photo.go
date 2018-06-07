package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type CreatePhotoRequest struct {
	TakenAt         int64  `position:"Query" name:"TakenAt"`
	PhotoTitle      string `position:"Query" name:"PhotoTitle"`
	LibraryId       string `position:"Query" name:"LibraryId"`
	ShareExpireTime int64  `position:"Query" name:"ShareExpireTime"`
	StoreName       string `position:"Query" name:"StoreName"`
	UploadType      string `position:"Query" name:"UploadType"`
	Remark          string `position:"Query" name:"Remark"`
	SessionId       string `position:"Query" name:"SessionId"`
	Staging         string `position:"Query" name:"Staging"`
	FileId          string `position:"Query" name:"FileId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreatePhotoRequest) Invoke(client goaliyun.Client) (*CreatePhotoResponse, error) {
	resp := &CreatePhotoResponse{}
	err := client.Request("cloudphoto", "CreatePhoto", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePhotoResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Photo     CreatePhotoPhoto
}

type CreatePhotoPhoto struct {
	Id              goaliyun.Integer
	IdStr           goaliyun.String
	Title           goaliyun.String
	FileId          goaliyun.String
	Location        goaliyun.String
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
