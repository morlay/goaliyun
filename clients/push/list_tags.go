package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagsRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListTagsRequest) Invoke(client goaliyun.Client) (*ListTagsResponse, error) {
	resp := &ListTagsResponse{}
	err := client.Request("push", "ListTags", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagsResponse struct {
	RequestId goaliyun.String
	TagInfos  ListTagsTagInfoList
}

type ListTagsTagInfo struct {
	TagName goaliyun.String
}

type ListTagsTagInfoList []ListTagsTagInfo

func (list *ListTagsTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagsTagInfo)
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
