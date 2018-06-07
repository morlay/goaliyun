package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAsrJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryAsrJobListRequest) Invoke(client goaliyun.Client) (*QueryAsrJobListResponse, error) {
	resp := &QueryAsrJobListResponse{}
	err := client.Request("mts", "QueryAsrJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAsrJobListResponse struct {
	RequestId   goaliyun.String
	JobList     QueryAsrJobListJobList
	NonExistIds QueryAsrJobListNonExistIdList
}

type QueryAsrJobListJob struct {
	Id           goaliyun.String
	UserData     goaliyun.String
	PipelineId   goaliyun.String
	State        goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Input        QueryAsrJobListInput
	AsrConfig    QueryAsrJobListAsrConfig
	AsrResult    QueryAsrJobListAsrResult
}

type QueryAsrJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryAsrJobListAsrConfig struct {
	Scene goaliyun.String
}

type QueryAsrJobListAsrResult struct {
	Duration    goaliyun.String
	AsrTextList QueryAsrJobListAsrTextList
}

type QueryAsrJobListAsrText struct {
	StartTime  goaliyun.Integer
	EndTime    goaliyun.String
	ChannelId  goaliyun.String
	SpeechRate goaliyun.String
	Text       goaliyun.String
}

type QueryAsrJobListJobList []QueryAsrJobListJob

func (list *QueryAsrJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrJobListJob)
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

type QueryAsrJobListNonExistIdList []goaliyun.String

func (list *QueryAsrJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryAsrJobListAsrTextList []QueryAsrJobListAsrText

func (list *QueryAsrJobListAsrTextList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrJobListAsrText)
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
