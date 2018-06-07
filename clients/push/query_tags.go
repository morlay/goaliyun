package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTagsRequest struct {
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QueryTagsRequest) Invoke(client goaliyun.Client) (*QueryTagsResponse, error) {
	resp := &QueryTagsResponse{}
	err := client.Request("push", "QueryTags", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTagsResponse struct {
	RequestId goaliyun.String
	TagInfos  QueryTagsTagInfoList
}

type QueryTagsTagInfo struct {
	TagName goaliyun.String
}

type QueryTagsTagInfoList []QueryTagsTagInfo

func (list *QueryTagsTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagsTagInfo)
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
