package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIASRJobRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AIASRJobIds          string `position:"Query" name:"AIASRJobIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAIASRJobRequest) Invoke(client goaliyun.Client) (*ListAIASRJobResponse, error) {
	resp := &ListAIASRJobResponse{}
	err := client.Request("vod", "ListAIASRJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIASRJobResponse struct {
	RequestId           goaliyun.String
	AIASRJobList        ListAIASRJobAIASRJobList
	NonExistAIASRJobIds ListAIASRJobNonExistAIASRJobIdList
}

type ListAIASRJobAIASRJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type ListAIASRJobAIASRJobList []ListAIASRJobAIASRJob

func (list *ListAIASRJobAIASRJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIASRJobAIASRJob)
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

type ListAIASRJobNonExistAIASRJobIdList []goaliyun.String

func (list *ListAIASRJobNonExistAIASRJobIdList) UnmarshalJSON(data []byte) error {
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
