package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIVideoCategoryJobRequest struct {
	ResourceOwnerId       string `position:"Query" name:"ResourceOwnerId"`
	AIVideoCategoryJobIds string `position:"Query" name:"AIVideoCategoryJobIds"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               string `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ListAIVideoCategoryJobRequest) Invoke(client goaliyun.Client) (*ListAIVideoCategoryJobResponse, error) {
	resp := &ListAIVideoCategoryJobResponse{}
	err := client.Request("vod", "ListAIVideoCategoryJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIVideoCategoryJobResponse struct {
	RequestId                     goaliyun.String
	AIVideoCategoryJobList        ListAIVideoCategoryJobAIVideoCategoryJobList
	NonExistAIVideoCategoryJobIds ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList
}

type ListAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type ListAIVideoCategoryJobAIVideoCategoryJobList []ListAIVideoCategoryJobAIVideoCategoryJob

func (list *ListAIVideoCategoryJobAIVideoCategoryJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCategoryJobAIVideoCategoryJob)
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

type ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList []goaliyun.String

func (list *ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList) UnmarshalJSON(data []byte) error {
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
