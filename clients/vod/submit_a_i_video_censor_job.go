package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitAIVideoCensorJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	AIVideoCensorConfig  string `position:"Query" name:"AIVideoCensorConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIVideoCensorJobRequest) Invoke(client goaliyun.Client) (*SubmitAIVideoCensorJobResponse, error) {
	resp := &SubmitAIVideoCensorJobResponse{}
	err := client.Request("vod", "SubmitAIVideoCensorJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIVideoCensorJobResponse struct {
	RequestId        goaliyun.String
	AIVideoCensorJob SubmitAIVideoCensorJobAIVideoCensorJob
}

type SubmitAIVideoCensorJobAIVideoCensorJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}
