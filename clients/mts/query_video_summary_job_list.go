package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryVideoSummaryJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryVideoSummaryJobListRequest) Invoke(client goaliyun.Client) (*QueryVideoSummaryJobListResponse, error) {
	resp := &QueryVideoSummaryJobListResponse{}
	err := client.Request("mts", "QueryVideoSummaryJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryVideoSummaryJobListResponse struct {
	RequestId   goaliyun.String
	JobList     QueryVideoSummaryJobListJobList
	NonExistIds QueryVideoSummaryJobListNonExistIdList
}

type QueryVideoSummaryJobListJob struct {
	Id                 goaliyun.String
	UserData           goaliyun.String
	PipelineId         goaliyun.String
	State              goaliyun.String
	Code               goaliyun.String
	Message            goaliyun.String
	CreationTime       goaliyun.String
	Input              QueryVideoSummaryJobListInput
	VideoSummaryResult QueryVideoSummaryJobListVideoSummaryResult
}

type QueryVideoSummaryJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryVideoSummaryJobListVideoSummaryResult struct {
	VideoSummaryList QueryVideoSummaryJobListVideoSummaryList
}

type QueryVideoSummaryJobListVideoSummary struct {
	StartTime goaliyun.String
	EndTime   goaliyun.String
}

type QueryVideoSummaryJobListJobList []QueryVideoSummaryJobListJob

func (list *QueryVideoSummaryJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryJobListJob)
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

type QueryVideoSummaryJobListNonExistIdList []goaliyun.String

func (list *QueryVideoSummaryJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryVideoSummaryJobListVideoSummaryList []QueryVideoSummaryJobListVideoSummary

func (list *QueryVideoSummaryJobListVideoSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryJobListVideoSummary)
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
