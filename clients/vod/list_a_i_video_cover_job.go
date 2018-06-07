package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAIVideoCoverJobRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AIVideoCoverJobIds   string `position:"Query" name:"AIVideoCoverJobIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListAIVideoCoverJobRequest) Invoke(client goaliyun.Client) (*ListAIVideoCoverJobResponse, error) {
	resp := &ListAIVideoCoverJobResponse{}
	err := client.Request("vod", "ListAIVideoCoverJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAIVideoCoverJobResponse struct {
	RequestId                  goaliyun.String
	AIVideoCoverJobList        ListAIVideoCoverJobAIVideoCoverJobList
	NonExistAIVideoCoverJobIds ListAIVideoCoverJobNonExistAIVideoCoverJobIdList
}

type ListAIVideoCoverJobAIVideoCoverJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type ListAIVideoCoverJobAIVideoCoverJobList []ListAIVideoCoverJobAIVideoCoverJob

func (list *ListAIVideoCoverJobAIVideoCoverJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCoverJobAIVideoCoverJob)
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

type ListAIVideoCoverJobNonExistAIVideoCoverJobIdList []goaliyun.String

func (list *ListAIVideoCoverJobNonExistAIVideoCoverJobIdList) UnmarshalJSON(data []byte) error {
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
