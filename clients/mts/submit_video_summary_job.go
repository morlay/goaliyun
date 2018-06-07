package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitVideoSummaryJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VideoSummaryConfig   string `position:"Query" name:"VideoSummaryConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitVideoSummaryJobRequest) Invoke(client goaliyun.Client) (*SubmitVideoSummaryJobResponse, error) {
	resp := &SubmitVideoSummaryJobResponse{}
	err := client.Request("mts", "SubmitVideoSummaryJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitVideoSummaryJobResponse struct {
	RequestId goaliyun.String
	JobId     goaliyun.String
}
