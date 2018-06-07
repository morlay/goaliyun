package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchMediaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	CateId               string `position:"Query" name:"CateId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	From                 string `position:"Query" name:"From"`
	SortBy               string `position:"Query" name:"SortBy"`
	To                   string `position:"Query" name:"To"`
	Tag                  string `position:"Query" name:"Tag"`
	KeyWord              string `position:"Query" name:"KeyWord"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SearchMediaRequest) Invoke(client goaliyun.Client) (*SearchMediaResponse, error) {
	resp := &SearchMediaResponse{}
	err := client.Request("mts", "SearchMedia", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchMediaResponse struct {
	RequestId  goaliyun.String
	TotalNum   goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	MediaList  SearchMediaMediaList
}

type SearchMediaMedia struct {
	MediaId      goaliyun.String
	Title        goaliyun.String
	Description  goaliyun.String
	CoverURL     goaliyun.String
	CateId       goaliyun.Integer
	Duration     goaliyun.String
	Format       goaliyun.String
	Size         goaliyun.String
	Bitrate      goaliyun.String
	Width        goaliyun.String
	Height       goaliyun.String
	Fps          goaliyun.String
	PublishState goaliyun.String
	CreationTime goaliyun.String
	Tags         SearchMediaTagList
	RunIdList    SearchMediaRunIdListList
	File         SearchMediaFile
}

type SearchMediaFile struct {
	URL   goaliyun.String
	State goaliyun.String
}

type SearchMediaMediaList []SearchMediaMedia

func (list *SearchMediaMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchMediaMedia)
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

type SearchMediaTagList []goaliyun.String

func (list *SearchMediaTagList) UnmarshalJSON(data []byte) error {
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

type SearchMediaRunIdListList []goaliyun.String

func (list *SearchMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
