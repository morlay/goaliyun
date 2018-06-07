package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitAIVideoTerrorismRecogJobRequest struct {
	UserData                    string `position:"Query" name:"UserData"`
	ResourceOwnerId             string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	AIVideoTerrorismRecogConfig string `position:"Query" name:"AIVideoTerrorismRecogConfig"`
	OwnerId                     string `position:"Query" name:"OwnerId"`
	MediaId                     string `position:"Query" name:"MediaId"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIVideoTerrorismRecogJobRequest) Invoke(client goaliyun.Client) (*SubmitAIVideoTerrorismRecogJobResponse, error) {
	resp := &SubmitAIVideoTerrorismRecogJobResponse{}
	err := client.Request("vod", "SubmitAIVideoTerrorismRecogJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIVideoTerrorismRecogJobResponse struct {
	RequestId                goaliyun.String
	AIVideoTerrorismRecogJob SubmitAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob
}

type SubmitAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}
