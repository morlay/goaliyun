package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIVideoCensorJobRequest struct {
	AIVideoCensorJobIds  string `position:"Query" name:"AIVideoCensorJobIds"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAIVideoCensorJobRequest) Invoke(client goaliyun.Client) (*ListAIVideoCensorJobResponse, error) {
	resp := &ListAIVideoCensorJobResponse{}
	err := client.Request("vod", "ListAIVideoCensorJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIVideoCensorJobResponse struct {
	RequestId                   goaliyun.String
	AIVideoCensorJobList        ListAIVideoCensorJobAIVideoCensorJobList
	NonExistAIVideoCensorJobIds ListAIVideoCensorJobNonExistAIVideoCensorJobIdList
}

type ListAIVideoCensorJobAIVideoCensorJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type ListAIVideoCensorJobAIVideoCensorJobList []ListAIVideoCensorJobAIVideoCensorJob

func (list *ListAIVideoCensorJobAIVideoCensorJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCensorJobAIVideoCensorJob)
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

type ListAIVideoCensorJobNonExistAIVideoCensorJobIdList []goaliyun.String

func (list *ListAIVideoCensorJobNonExistAIVideoCensorJobIdList) UnmarshalJSON(data []byte) error {
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
