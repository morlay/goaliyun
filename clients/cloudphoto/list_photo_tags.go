package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPhotoTagsRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	Lang      string `position:"Query" name:"Lang"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListPhotoTagsRequest) Invoke(client goaliyun.Client) (*ListPhotoTagsResponse, error) {
	resp := &ListPhotoTagsResponse{}
	err := client.Request("cloudphoto", "ListPhotoTags", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPhotoTagsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Tags      ListPhotoTagsTagList
}

type ListPhotoTagsTag struct {
	Id        goaliyun.Integer
	IdStr     goaliyun.String
	IsSubTag  bool
	Name      goaliyun.String
	ParentTag goaliyun.String
}

type ListPhotoTagsTagList []ListPhotoTagsTag

func (list *ListPhotoTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoTagsTag)
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
