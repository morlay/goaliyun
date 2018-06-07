package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagsRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Lang      string `position:"Query" name:"Lang"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListTagsRequest) Invoke(client goaliyun.Client) (*ListTagsResponse, error) {
	resp := &ListTagsResponse{}
	err := client.Request("cloudphoto", "ListTags", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagsResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
	Tags      ListTagsTagList
}

type ListTagsTag struct {
	Id        goaliyun.Integer
	IdStr     goaliyun.String
	Name      goaliyun.String
	IsSubTag  bool
	ParentTag goaliyun.String
	Cover     ListTagsCover
}

type ListTagsCover struct {
	Id      goaliyun.Integer
	IdStr   goaliyun.String
	Title   goaliyun.String
	FileId  goaliyun.String
	State   goaliyun.String
	Md5     goaliyun.String
	IsVideo bool
	Remark  goaliyun.String
	Width   goaliyun.Integer
	Height  goaliyun.Integer
	Ctime   goaliyun.Integer
	Mtime   goaliyun.Integer
}

type ListTagsTagList []ListTagsTag

func (list *ListTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagsTag)
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
