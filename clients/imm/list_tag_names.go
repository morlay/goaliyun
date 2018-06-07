package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagNamesRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListTagNamesRequest) Invoke(client goaliyun.Client) (*ListTagNamesResponse, error) {
	resp := &ListTagNamesResponse{}
	err := client.Request("imm", "ListTagNames", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagNamesResponse struct {
	RequestId goaliyun.String
	Tags      ListTagNamesTagsItemList
}

type ListTagNamesTagsItem struct {
	TagName goaliyun.String
	Num     goaliyun.Integer
}

type ListTagNamesTagsItemList []ListTagNamesTagsItem

func (list *ListTagNamesTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagNamesTagsItem)
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
