package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIVideoSummaryJobRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AIVideoSummaryJobIds string `position:"Query" name:"AIVideoSummaryJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAIVideoSummaryJobRequest) Invoke(client goaliyun.Client) (*ListAIVideoSummaryJobResponse, error) {
	resp := &ListAIVideoSummaryJobResponse{}
	err := client.Request("vod", "ListAIVideoSummaryJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIVideoSummaryJobResponse struct {
	RequestId                    goaliyun.String
	AIVideoSummaryJobList        ListAIVideoSummaryJobAIVideoSummaryJobList
	NonExistAIVideoSummaryJobIds ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList
}

type ListAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type ListAIVideoSummaryJobAIVideoSummaryJobList []ListAIVideoSummaryJobAIVideoSummaryJob

func (list *ListAIVideoSummaryJobAIVideoSummaryJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoSummaryJobAIVideoSummaryJob)
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

type ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList []goaliyun.String

func (list *ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList) UnmarshalJSON(data []byte) error {
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
