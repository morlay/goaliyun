package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitAIVideoSummaryJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	AIVideoSummaryConfig string `position:"Query" name:"AIVideoSummaryConfig"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIVideoSummaryJobRequest) Invoke(client goaliyun.Client) (*SubmitAIVideoSummaryJobResponse, error) {
	resp := &SubmitAIVideoSummaryJobResponse{}
	err := client.Request("vod", "SubmitAIVideoSummaryJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIVideoSummaryJobResponse struct {
	RequestId         goaliyun.String
	AIVideoSummaryJob SubmitAIVideoSummaryJobAIVideoSummaryJob
}

type SubmitAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}
