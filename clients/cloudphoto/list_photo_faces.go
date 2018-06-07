package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPhotoFacesRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListPhotoFacesRequest) Invoke(client goaliyun.Client) (*ListPhotoFacesResponse, error) {
	resp := &ListPhotoFacesResponse{}
	err := client.Request("cloudphoto", "ListPhotoFaces", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPhotoFacesResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Faces     ListPhotoFacesFaceList
}

type ListPhotoFacesFace struct {
	FaceId    goaliyun.Integer
	FaceIdStr goaliyun.String
	FaceName  goaliyun.String
	Axis      ListPhotoFacesAxiList
}

type ListPhotoFacesFaceList []ListPhotoFacesFace

func (list *ListPhotoFacesFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoFacesFace)
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

type ListPhotoFacesAxiList []goaliyun.String

func (list *ListPhotoFacesAxiList) UnmarshalJSON(data []byte) error {
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
