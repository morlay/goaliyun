package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIJobRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAIJobRequest) Invoke(client goaliyun.Client) (*ListAIJobResponse, error) {
	resp := &ListAIJobResponse{}
	err := client.Request("vod", "ListAIJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIJobResponse struct {
	RequestId        goaliyun.String
	AIJobList        ListAIJobAIJobList
	NonExistAIJobIds ListAIJobNonExistAIJobIdList
}

type ListAIJobAIJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Type         goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	CompleteTime goaliyun.String
	Data         goaliyun.String
}

type ListAIJobAIJobList []ListAIJobAIJob

func (list *ListAIJobAIJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIJobAIJob)
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

type ListAIJobNonExistAIJobIdList []goaliyun.String

func (list *ListAIJobNonExistAIJobIdList) UnmarshalJSON(data []byte) error {
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
