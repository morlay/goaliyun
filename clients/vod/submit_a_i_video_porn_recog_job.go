package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitAIVideoPornRecogJobRequest struct {
	UserData               string `position:"Query" name:"UserData"`
	ResourceOwnerId        string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	AIVideoPornRecogConfig string `position:"Query" name:"AIVideoPornRecogConfig"`
	OwnerId                string `position:"Query" name:"OwnerId"`
	MediaId                string `position:"Query" name:"MediaId"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIVideoPornRecogJobRequest) Invoke(client goaliyun.Client) (*SubmitAIVideoPornRecogJobResponse, error) {
	resp := &SubmitAIVideoPornRecogJobResponse{}
	err := client.Request("vod", "SubmitAIVideoPornRecogJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIVideoPornRecogJobResponse struct {
	RequestId           goaliyun.String
	AIVideoPornRecogJob SubmitAIVideoPornRecogJobAIVideoPornRecogJob
}

type SubmitAIVideoPornRecogJobAIVideoPornRecogJob struct {
	JobId        goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}
