package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitAIASRJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AIASRConfig          string `position:"Query" name:"AIASRConfig"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIASRJobRequest) Invoke(client goaliyun.Client) (*SubmitAIASRJobResponse, error) {
	resp := &SubmitAIASRJobResponse{}
	err := client.Request("vod", "SubmitAIASRJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIASRJobResponse struct {
	RequestId goaliyun.String
	AIASRJob  SubmitAIASRJobAIASRJob
}

type SubmitAIASRJobAIASRJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}
