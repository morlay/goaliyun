package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagNameRequest struct {
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListTagNameRequest) Invoke(client goaliyun.Client) (*ListTagNameResponse, error) {
	resp := &ListTagNameResponse{}
	err := client.Request("imm", "ListTagName", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagNameResponse struct {
	RequestId goaliyun.String
	Tags      ListTagNameTagsItemList
}

type ListTagNameTagsItem struct {
	TagName goaliyun.String
	Num     goaliyun.Integer
}

type ListTagNameTagsItemList []ListTagNameTagsItem

func (list *ListTagNameTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagNameTagsItem)
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
