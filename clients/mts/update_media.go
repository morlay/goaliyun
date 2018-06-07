package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpdateMediaRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Title                string `position:"Query" name:"Title"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateMediaRequest) Invoke(client goaliyun.Client) (*UpdateMediaResponse, error) {
	resp := &UpdateMediaResponse{}
	err := client.Request("mts", "UpdateMedia", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateMediaResponse struct {
	RequestId goaliyun.String
	Media     UpdateMediaMedia
}

type UpdateMediaMedia struct {
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
	Tags         UpdateMediaTagList
	RunIdList    UpdateMediaRunIdListList
	File         UpdateMediaFile
}

type UpdateMediaFile struct {
	URL   goaliyun.String
	State goaliyun.String
}

type UpdateMediaTagList []goaliyun.String

func (list *UpdateMediaTagList) UnmarshalJSON(data []byte) error {
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

type UpdateMediaRunIdListList []goaliyun.String

func (list *UpdateMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
