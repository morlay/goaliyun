package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryVideoSplitJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryVideoSplitJobListRequest) Invoke(client goaliyun.Client) (*QueryVideoSplitJobListResponse, error) {
	resp := &QueryVideoSplitJobListResponse{}
	err := client.Request("mts", "QueryVideoSplitJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryVideoSplitJobListResponse struct {
	RequestId   goaliyun.String
	JobList     QueryVideoSplitJobListJobList
	NonExistIds QueryVideoSplitJobListNonExistIdList
}

type QueryVideoSplitJobListJob struct {
	Id               goaliyun.String
	UserData         goaliyun.String
	PipelineId       goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	CreationTime     goaliyun.String
	Input            QueryVideoSplitJobListInput
	VideoSplitResult QueryVideoSplitJobListVideoSplitResult
}

type QueryVideoSplitJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryVideoSplitJobListVideoSplitResult struct {
	VideoSplitList QueryVideoSplitJobListVideoSplitList
}

type QueryVideoSplitJobListVideoSplit struct {
	StartTime goaliyun.String
	EndTime   goaliyun.String
	Path      goaliyun.String
}

type QueryVideoSplitJobListJobList []QueryVideoSplitJobListJob

func (list *QueryVideoSplitJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSplitJobListJob)
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

type QueryVideoSplitJobListNonExistIdList []goaliyun.String

func (list *QueryVideoSplitJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryVideoSplitJobListVideoSplitList []QueryVideoSplitJobListVideoSplit

func (list *QueryVideoSplitJobListVideoSplitList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSplitJobListVideoSplit)
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
