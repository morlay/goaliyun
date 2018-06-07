package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitAIVideoCoverJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	AIVideoCoverConfig   string `position:"Query" name:"AIVideoCoverConfig"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIVideoCoverJobRequest) Invoke(client goaliyun.Client) (*SubmitAIVideoCoverJobResponse, error) {
	resp := &SubmitAIVideoCoverJobResponse{}
	err := client.Request("vod", "SubmitAIVideoCoverJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIVideoCoverJobResponse struct {
	RequestId       goaliyun.String
	AIVideoCoverJob SubmitAIVideoCoverJobAIVideoCoverJob
}

type SubmitAIVideoCoverJobAIVideoCoverJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}
