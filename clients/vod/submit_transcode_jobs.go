package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitTranscodeJobsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TemplateGroupId      string `position:"Query" name:"TemplateGroupId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	EncryptConfig        string `position:"Query" name:"EncryptConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitTranscodeJobsRequest) Invoke(client goaliyun.Client) (*SubmitTranscodeJobsResponse, error) {
	resp := &SubmitTranscodeJobsResponse{}
	err := client.Request("vod", "SubmitTranscodeJobs", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitTranscodeJobsResponse struct {
	RequestId     goaliyun.String
	TranscodeJobs SubmitTranscodeJobsTranscodeJobList
}

type SubmitTranscodeJobsTranscodeJob struct {
	JobId goaliyun.String
}

type SubmitTranscodeJobsTranscodeJobList []SubmitTranscodeJobsTranscodeJob

func (list *SubmitTranscodeJobsTranscodeJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitTranscodeJobsTranscodeJob)
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
