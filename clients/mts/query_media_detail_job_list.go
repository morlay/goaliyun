package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMediaDetailJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryMediaDetailJobListRequest) Invoke(client goaliyun.Client) (*QueryMediaDetailJobListResponse, error) {
	resp := &QueryMediaDetailJobListResponse{}
	err := client.Request("mts", "QueryMediaDetailJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMediaDetailJobListResponse struct {
	RequestId   goaliyun.String
	JobList     QueryMediaDetailJobListJobList
	NonExistIds QueryMediaDetailJobListNonExistIdList
}

type QueryMediaDetailJobListJob struct {
	Id                goaliyun.String
	UserData          goaliyun.String
	PipelineId        goaliyun.String
	State             goaliyun.String
	Code              goaliyun.String
	Message           goaliyun.String
	CreationTime      goaliyun.String
	Input             QueryMediaDetailJobListInput
	MediaDetailConfig QueryMediaDetailJobListMediaDetailConfig
	MediaDetailResult QueryMediaDetailJobListMediaDetailResult
}

type QueryMediaDetailJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryMediaDetailJobListMediaDetailConfig struct {
	Scenario   goaliyun.String
	DetailType goaliyun.String
	OutputFile QueryMediaDetailJobListOutputFile
}

type QueryMediaDetailJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryMediaDetailJobListMediaDetailResult struct {
	Status                 goaliyun.String
	MediaDetailRecgResults QueryMediaDetailJobListMediaDetailRecgResultList
	Tags                   QueryMediaDetailJobListTagList
}

type QueryMediaDetailJobListMediaDetailRecgResult struct {
	ImageUrl      goaliyun.String
	Time          goaliyun.String
	OcrText       goaliyun.String
	Celebrities   QueryMediaDetailJobListCelebrityList
	Sensitives    QueryMediaDetailJobListSensitiveList
	Politicians   QueryMediaDetailJobListPoliticianList
	FrameTagInfos QueryMediaDetailJobListFrameTagInfoList
	FrameTags     QueryMediaDetailJobListFrameTagList
}

type QueryMediaDetailJobListCelebrity struct {
	Name   goaliyun.String
	Score  goaliyun.String
	Target goaliyun.String
}

type QueryMediaDetailJobListSensitive struct {
	Name   goaliyun.String
	Score  goaliyun.String
	Target goaliyun.String
}

type QueryMediaDetailJobListPolitician struct {
	Name   goaliyun.String
	Score  goaliyun.String
	Target goaliyun.String
}

type QueryMediaDetailJobListFrameTagInfo struct {
	Tag      goaliyun.String
	Score    goaliyun.String
	Category goaliyun.String
}

type QueryMediaDetailJobListJobList []QueryMediaDetailJobListJob

func (list *QueryMediaDetailJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListJob)
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

type QueryMediaDetailJobListNonExistIdList []goaliyun.String

func (list *QueryMediaDetailJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaDetailJobListMediaDetailRecgResultList []QueryMediaDetailJobListMediaDetailRecgResult

func (list *QueryMediaDetailJobListMediaDetailRecgResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListMediaDetailRecgResult)
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

type QueryMediaDetailJobListTagList []goaliyun.String

func (list *QueryMediaDetailJobListTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaDetailJobListCelebrityList []QueryMediaDetailJobListCelebrity

func (list *QueryMediaDetailJobListCelebrityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListCelebrity)
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

type QueryMediaDetailJobListSensitiveList []QueryMediaDetailJobListSensitive

func (list *QueryMediaDetailJobListSensitiveList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListSensitive)
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

type QueryMediaDetailJobListPoliticianList []QueryMediaDetailJobListPolitician

func (list *QueryMediaDetailJobListPoliticianList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListPolitician)
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

type QueryMediaDetailJobListFrameTagInfoList []QueryMediaDetailJobListFrameTagInfo

func (list *QueryMediaDetailJobListFrameTagInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaDetailJobListFrameTagInfo)
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

type QueryMediaDetailJobListFrameTagList []goaliyun.String

func (list *QueryMediaDetailJobListFrameTagList) UnmarshalJSON(data []byte) error {
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
