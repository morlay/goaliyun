package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRegisteredTagsRequest struct {
	StoreName string                      `position:"Query" name:"StoreName"`
	Langs     *ListRegisteredTagsLangList `position:"Query" type:"Repeated" name:"Lang"`
	RegionId  string                      `position:"Query" name:"RegionId"`
}

func (req *ListRegisteredTagsRequest) Invoke(client goaliyun.Client) (*ListRegisteredTagsResponse, error) {
	resp := &ListRegisteredTagsResponse{}
	err := client.Request("cloudphoto", "ListRegisteredTags", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRegisteredTagsResponse struct {
	Code           goaliyun.String
	Message        goaliyun.String
	RequestId      goaliyun.String
	Action         goaliyun.String
	RegisteredTags ListRegisteredTagsRegisteredTagList
}

type ListRegisteredTagsRegisteredTag struct {
	TagKey    goaliyun.String
	TagValues ListRegisteredTagsTagValueList
}

type ListRegisteredTagsTagValue struct {
	Lang goaliyun.String
	Text goaliyun.String
}

type ListRegisteredTagsLangList []string

func (list *ListRegisteredTagsLangList) UnmarshalJSON(data []byte) error {
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

type ListRegisteredTagsRegisteredTagList []ListRegisteredTagsRegisteredTag

func (list *ListRegisteredTagsRegisteredTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegisteredTagsRegisteredTag)
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

type ListRegisteredTagsTagValueList []ListRegisteredTagsTagValue

func (list *ListRegisteredTagsTagValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRegisteredTagsTagValue)
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
