package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIVideoPornRecogJobRequest struct {
	ResourceOwnerId        string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	AIVideoPornRecogJobIds string `position:"Query" name:"AIVideoPornRecogJobIds"`
	OwnerId                string `position:"Query" name:"OwnerId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *ListAIVideoPornRecogJobRequest) Invoke(client goaliyun.Client) (*ListAIVideoPornRecogJobResponse, error) {
	resp := &ListAIVideoPornRecogJobResponse{}
	err := client.Request("vod", "ListAIVideoPornRecogJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIVideoPornRecogJobResponse struct {
	RequestId               goaliyun.String
	AIVideoPornRecogJobList ListAIVideoPornRecogJobAIVideoPornRecogJobList
	NonExistPornRecogJobIds ListAIVideoPornRecogJobNonExistPornRecogJobIdList
}

type ListAIVideoPornRecogJobAIVideoPornRecogJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type ListAIVideoPornRecogJobAIVideoPornRecogJobList []ListAIVideoPornRecogJobAIVideoPornRecogJob

func (list *ListAIVideoPornRecogJobAIVideoPornRecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoPornRecogJobAIVideoPornRecogJob)
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

type ListAIVideoPornRecogJobNonExistPornRecogJobIdList []goaliyun.String

func (list *ListAIVideoPornRecogJobNonExistPornRecogJobIdList) UnmarshalJSON(data []byte) error {
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
