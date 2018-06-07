package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIVideoTerrorismRecogJobRequest struct {
	ResourceOwnerId             string `position:"Query" name:"ResourceOwnerId"`
	AIVideoTerrorismRecogJobIds string `position:"Query" name:"AIVideoTerrorismRecogJobIds"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	OwnerId                     string `position:"Query" name:"OwnerId"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *ListAIVideoTerrorismRecogJobRequest) Invoke(client goaliyun.Client) (*ListAIVideoTerrorismRecogJobResponse, error) {
	resp := &ListAIVideoTerrorismRecogJobResponse{}
	err := client.Request("vod", "ListAIVideoTerrorismRecogJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIVideoTerrorismRecogJobResponse struct {
	RequestId                    goaliyun.String
	AIVideoTerrorismRecogJobList ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList
	NonExistTerrorismRecogJobIds ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList []ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob

func (list *ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob)
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

type ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList []goaliyun.String

func (list *ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList) UnmarshalJSON(data []byte) error {
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
