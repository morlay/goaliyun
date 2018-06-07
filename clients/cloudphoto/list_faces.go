package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFacesRequest struct {
	Cursor      string `position:"Query" name:"Cursor"`
	HasFaceName string `position:"Query" name:"HasFaceName"`
	Size        int64  `position:"Query" name:"Size"`
	LibraryId   string `position:"Query" name:"LibraryId"`
	StoreName   string `position:"Query" name:"StoreName"`
	State       string `position:"Query" name:"State"`
	Direction   string `position:"Query" name:"Direction"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ListFacesRequest) Invoke(client goaliyun.Client) (*ListFacesResponse, error) {
	resp := &ListFacesResponse{}
	err := client.Request("cloudphoto", "ListFaces", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFacesResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	NextCursor goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Faces      ListFacesFaceList
}

type ListFacesFace struct {
	Id          goaliyun.Integer
	IdStr       goaliyun.String
	Name        goaliyun.String
	PhotosCount goaliyun.Integer
	State       goaliyun.String
	IsMe        bool
	Ctime       goaliyun.Integer
	Mtime       goaliyun.Integer
	Axis        ListFacesAxiList
	Cover       ListFacesCover
}

type ListFacesCover struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Title   goaliyun.String
	FileId  goaliyun.String
	State   goaliyun.String
	Md5     goaliyun.String
	IsVideo bool
	Width   goaliyun.Integer
	Height  goaliyun.Integer
	Ctime   goaliyun.Integer
	Mtime   goaliyun.Integer
	Remark  goaliyun.String
}

type ListFacesFaceList []ListFacesFace

func (list *ListFacesFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFacesFace)
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

type ListFacesAxiList []goaliyun.String

func (list *ListFacesAxiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
