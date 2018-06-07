package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMediaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	NextPageToken        string `position:"Query" name:"NextPageToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      int64  `position:"Query" name:"MaximumPageSize"`
	From                 string `position:"Query" name:"From"`
	To                   string `position:"Query" name:"To"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListMediaRequest) Invoke(client goaliyun.Client) (*ListMediaResponse, error) {
	resp := &ListMediaResponse{}
	err := client.Request("mts", "ListMedia", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMediaResponse struct {
	RequestId     goaliyun.String
	NextPageToken goaliyun.String
	MediaList     ListMediaMediaList
}

type ListMediaMedia struct {
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
	Tags         ListMediaTagList
	RunIdList    ListMediaRunIdListList
	File         ListMediaFile
}

type ListMediaFile struct {
	URL   goaliyun.String
	State goaliyun.String
}

type ListMediaMediaList []ListMediaMedia

func (list *ListMediaMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaMedia)
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

type ListMediaTagList []goaliyun.String

func (list *ListMediaTagList) UnmarshalJSON(data []byte) error {
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

type ListMediaRunIdListList []goaliyun.String

func (list *ListMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
