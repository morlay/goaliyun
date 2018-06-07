package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetPhotosByMd5sRequest struct {
	LibraryId string                  `position:"Query" name:"LibraryId"`
	StoreName string                  `position:"Query" name:"StoreName"`
	State     string                  `position:"Query" name:"State"`
	Md5s      *GetPhotosByMd5sMd5List `position:"Query" type:"Repeated" name:"Md.5"`
	RegionId  string                  `position:"Query" name:"RegionId"`
}

func (req *GetPhotosByMd5sRequest) Invoke(client goaliyun.Client) (*GetPhotosByMd5sResponse, error) {
	resp := &GetPhotosByMd5sResponse{}
	err := client.Request("cloudphoto", "GetPhotosByMd5s", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPhotosByMd5sResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Photos    GetPhotosByMd5sPhotoList
}

type GetPhotosByMd5sPhoto struct {
	Id              goaliyun.Integer
	IdStr           goaliyun.String
	Title           goaliyun.String
	FileId          goaliyun.String
	Location        goaliyun.String
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
}

type GetPhotosByMd5sMd5List []string

func (list *GetPhotosByMd5sMd5List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type GetPhotosByMd5sPhotoList []GetPhotosByMd5sPhoto

func (list *GetPhotosByMd5sPhotoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPhotosByMd5sPhoto)
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
