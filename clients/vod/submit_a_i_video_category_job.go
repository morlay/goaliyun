package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitAIVideoCategoryJobRequest struct {
	AIVideoCategoryConfig string `position:"Query" name:"AIVideoCategoryConfig"`
	UserData              string `position:"Query" name:"UserData"`
	ResourceOwnerId       string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               string `position:"Query" name:"OwnerId"`
	MediaId               string `position:"Query" name:"MediaId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIVideoCategoryJobRequest) Invoke(client goaliyun.Client) (*SubmitAIVideoCategoryJobResponse, error) {
	resp := &SubmitAIVideoCategoryJobResponse{}
	err := client.Request("vod", "SubmitAIVideoCategoryJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIVideoCategoryJobResponse struct {
	RequestId          goaliyun.String
	AIVideoCategoryJob SubmitAIVideoCategoryJobAIVideoCategoryJob
}

type SubmitAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}
