package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitAIJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Types                string `position:"Query" name:"Types"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Config               string `position:"Query" name:"Config"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitAIJobRequest) Invoke(client goaliyun.Client) (*SubmitAIJobResponse, error) {
	resp := &SubmitAIJobResponse{}
	err := client.Request("vod", "SubmitAIJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitAIJobResponse struct {
	RequestId goaliyun.String
	AIJobList SubmitAIJobAIJobList
}

type SubmitAIJobAIJob struct {
	JobId        goaliyun.String
	Type         goaliyun.String
	MediaId      goaliyun.String
	Status       goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	Data         goaliyun.String
}

type SubmitAIJobAIJobList []SubmitAIJobAIJob

func (list *SubmitAIJobAIJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitAIJobAIJob)
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
