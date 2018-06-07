package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddMediaRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	Description           string `position:"Query" name:"Description"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	Title                 string `position:"Query" name:"Title"`
	Tags                  string `position:"Query" name:"Tags"`
	CoverURL              string `position:"Query" name:"CoverURL"`
	CateId                int64  `position:"Query" name:"CateId"`
	FileURL               string `position:"Query" name:"FileURL"`
	MediaWorkflowId       string `position:"Query" name:"MediaWorkflowId"`
	MediaWorkflowUserData string `position:"Query" name:"MediaWorkflowUserData"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *AddMediaRequest) Invoke(client goaliyun.Client) (*AddMediaResponse, error) {
	resp := &AddMediaResponse{}
	err := client.Request("mts", "AddMedia", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddMediaResponse struct {
	RequestId goaliyun.String
	Media     AddMediaMedia
}

type AddMediaMedia struct {
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
	Tags         AddMediaTagList
	RunIdList    AddMediaRunIdListList
	File         AddMediaFile
}

type AddMediaFile struct {
	URL   goaliyun.String
	State goaliyun.String
}

type AddMediaTagList []goaliyun.String

func (list *AddMediaTagList) UnmarshalJSON(data []byte) error {
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

type AddMediaRunIdListList []goaliyun.String

func (list *AddMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
