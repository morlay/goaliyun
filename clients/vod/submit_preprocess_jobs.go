package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SubmitPreprocessJobsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PreprocessType       string `position:"Query" name:"PreprocessType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitPreprocessJobsRequest) Invoke(client goaliyun.Client) (*SubmitPreprocessJobsResponse, error) {
	resp := &SubmitPreprocessJobsResponse{}
	err := client.Request("vod", "SubmitPreprocessJobs", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitPreprocessJobsResponse struct {
	RequestId      goaliyun.String
	PreprocessJobs SubmitPreprocessJobsPreprocessJobList
}

type SubmitPreprocessJobsPreprocessJob struct {
	JobId goaliyun.String
}

type SubmitPreprocessJobsPreprocessJobList []SubmitPreprocessJobsPreprocessJob

func (list *SubmitPreprocessJobsPreprocessJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitPreprocessJobsPreprocessJob)
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
